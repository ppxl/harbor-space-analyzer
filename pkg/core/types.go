package core

import (
	artiModel "github.com/goharbor/harbor/src/pkg/artifact"
	projModel "github.com/goharbor/harbor/src/pkg/project/models"
	repoModel "github.com/goharbor/harbor/src/pkg/repository/model"
)

type AnalyzerArgs struct {
	Credentials *HarborCredentials
	HarborURL   string
}

type HarborCredentials struct {
	Username string
	Password string
}

type RepoSummary struct {
	RepoName      string
	ArtifactCount int64
	Size          int64
}

type ProjectSummary struct {
	ProjectName   string
	ArtifactCount int64
	Size          int64
	RepoSummaries []RepoSummary
}

type Artifact struct {
	artiModel.Artifact `json:",inline"`
}

type CountingRepository struct {
	repoModel.RepoRecord `json:",inline"`
	ArtifactCount        int `json:"artifact_count"`
}

type Project struct {
	projModel.Project `json:",inline"`
}
