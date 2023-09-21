package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ppxl/harbor-space-analyzer/pkg/core"
)

type AnalyzeService struct {
	args core.AnalyzerArgs
}

func New(args core.AnalyzerArgs) *AnalyzeService {
	return &AnalyzeService{args: args}
}

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
		println("next char", karacter, "=", prjSum.ProjectName)

		usagePercent = append(usagePercent, prjPercent)
		karacters = append(karacters, karacter)
	}

	// fix hole in diagram
	usagePercent[len(usagePercent)-1] += 0.001

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
	queryEndpoint := fmt.Sprintf("/projects/%s/repositories/%s/artifacts?page=1&page_size=100&with_tag=true'", projectName, repoName)
	response, err := as.createAuthdGetRequest(ctx, queryEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to request artifact for project %s repo %s: %w", projectName, repoName, err)
	}

	body, err := io.ReadAll(response.Body)
	var artifacts []core.Artifact
	err = json.Unmarshal(body, &artifacts)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal artifact for project %s repo %s: %s: %w", projectName, repoName, body, err)
	}

	return artifacts, nil
}

func (as *AnalyzeService) getProjects(ctx context.Context) ([]core.Project, error) {
	response, err := as.createAuthdGetRequest(ctx, "/projects?page=1&page_size=100&with_detail=false'")
	if err != nil {
		return nil, fmt.Errorf("failed to request projects: %w", err)
	}

	body, err := io.ReadAll(response.Body)
	var projects []core.Project
	err = json.Unmarshal(body, &projects)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects %s: %w", body, err)
	}

	return projects, nil
}

func (as *AnalyzeService) getReposByProject(ctx context.Context, prj core.Project) ([]core.CountingRepository, error) {
	response, err := as.createAuthdGetRequest(ctx, "/projects/"+prj.Name+"/repositories?page=1&page_size=100")
	if err != nil {
		return nil, fmt.Errorf("repo request failed for endpoint /projects/%s/repositories: %w", prj.Name, err)
	}

	body, err := io.ReadAll(response.Body)
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
	baseURL := as.args.HarborURL

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", "Basic "+basicAuth(as.args.Credentials.Username, as.args.Credentials.Password))
	request.Header.Add("accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	// debug log
	//fmt.Println("GET", baseURL+endpoint, "result:", response.Body)
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
