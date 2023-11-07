package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ppxl/harbor-space-analyzer/pkg/core"
)

const harborApiVersion = "/api/v2.0"

type AnalyzeService struct {
	args core.AnalyzerArgs
}

func New(args core.AnalyzerArgs) *AnalyzeService {
	return &AnalyzeService{args: args}
}

// CalculateValues calculates the percentages per character key and returns both.
func CalculateValues(summaries []core.ProjectSummary) (karacters []string, usagePercent []float64) {
	// maybe move to type alias of []ProjectSummary
	var usageTotal int64 = 0
	for _, prjSum := range summaries {
		usageTotal += prjSum.Size
	}

	if usageTotal == 0 {
		return
	}

	for idx, prjSum := range summaries {
		if prjSum.Size == 0 {
			continue
		}
		prjPercent := float64(prjSum.Size) / float64(usageTotal)
		karacter := string('A' + int32(idx))

		usagePercent = append(usagePercent, prjPercent)
		karacters = append(karacters, karacter)
	}

	return
}

func (as *AnalyzeService) GetProjectInfo(ctx context.Context) ([]core.ProjectSummary, error) {
	projects, err := as.getProjects(ctx)
	if err != nil {
		return nil, err
	}

	var projSums []core.ProjectSummary

	for _, proj := range projects {
		repos, err := as.getReposByProject(ctx, proj)
		if err != nil {
			return nil, err
		}

		repoSums, err := as.getRepoSummaries(ctx, proj.Name, repos)
		if err != nil {
			return nil, err
		}

		projectSum := core.ProjectSummary{ProjectName: proj.Name, Size: 0}
		for _, repoSum := range repoSums {
			projectSum.RepoSummaries = repoSums
			projectSum.ArtifactCount += repoSum.ArtifactCount
			projectSum.Size += repoSum.Size
		}

		projSums = append(projSums, projectSum)
	}

	return projSums, nil
}

func (as *AnalyzeService) getRepoSummaries(ctx context.Context, projectName string, repos []core.CountingRepository) ([]core.RepoSummary, error) {
	var repoSums []core.RepoSummary

	for _, repo := range repos {
		repoName := strings.Replace(repo.Name, projectName+"/", "", 1)
		fmt.Println("Looking at", projectName, repoName)
		artifacts, err := as.getArtifactsForRepo(ctx, projectName, repoName)
		if err != nil {
			return nil, err
		}

		repoSum := core.RepoSummary{RepoName: repo.Name, ArtifactCount: 0, Size: 0}
		for _, arti := range artifacts {
			repoSum.ArtifactCount++
			repoSum.Size += arti.Size
		}

		repoSums = append(repoSums, repoSum)
	}

	return repoSums, nil
}

func (as *AnalyzeService) getArtifactsForRepo(ctx context.Context, projectName string, repoName string) ([]core.Artifact, error) {
	repoName = escapeSlashes(repoName)
	endpoint := fmt.Sprintf("/projects/%s/repositories/%s/artifacts?page=1&page_size=100&with_tag=true", projectName, repoName)
	response, err := as.createAuthdGetRequest(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to request artifact for project %s repo %s: %w", projectName, repoName, err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body of request to %s: %w", endpoint, err)
	}
	var artifacts []core.Artifact
	err = json.Unmarshal(body, &artifacts)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal artifact for project %s repo %s: %s: %w", projectName, repoName, body, err)
	}

	return artifacts, nil
}

func (as *AnalyzeService) getProjects(ctx context.Context) ([]core.Project, error) {
	const endpoint = "/projects?page=1&page_size=100&with_detail=false"
	response, err := as.createAuthdGetRequest(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to request projects: %w", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body of request to %s: %w", endpoint, err)
	}
	var projects []core.Project
	err = json.Unmarshal(body, &projects)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal projects %s: %w", body, err)
	}

	return projects, nil
}

func (as *AnalyzeService) getReposByProject(ctx context.Context, prj core.Project) ([]core.CountingRepository, error) {
	endpoint := "/projects/" + prj.Name + "/repositories?page=1&page_size=100"
	response, err := as.createAuthdGetRequest(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("repo request failed for endpoint /projects/%s/repositories: %w", prj.Name, err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body of request to %s: %w", endpoint, err)
	}
	var repoList []core.CountingRepository
	err = json.Unmarshal(body, &repoList)
	if err != nil {
		return nil, fmt.Errorf("failed to to unmarshal %s repo data %s: %w", prj.Name, string(body), err)
	}

	// debug
	//fmt.Printf("Found %s repo counts:", prj.Name)
	//for _, repo := range repoList {
	//	fmt.Printf(" %s: %d, ", repo.Name, repo.ArtifactCount)
	//}
	//fmt.Println()

	return repoList, nil
}

func (as *AnalyzeService) createAuthdGetRequest(ctx context.Context, endpoint string) (*http.Response, error) {
	theUrl, err := url.JoinPath(as.args.HarborURL, harborApiVersion)
	if err != nil {
		return nil, fmt.Errorf("could not parse harbor URL with these fragments: %s, %s, %s: %w", as.args.HarborURL, harborApiVersion, err)
	}
	theUrl += endpoint

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, theUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", "Basic "+basicAuth(as.args.Credentials.Username, as.args.Credentials.Password))
	request.Header.Add("accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("request to theUrl failed: %w", err)
	}

	// debug log
	//fmt.Println("GET", theUrl, "result:", response.Body)
	return response, nil
}

func escapeSlashes(name string) string {
	// this should only be the case when harbor returns repositories with multiple slashes like project/repo/thing.
	return strings.ReplaceAll(name, "/", "%252F")
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
