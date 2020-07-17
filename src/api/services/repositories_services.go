package services

import (
	"golang-microservices/src/api/config"
	"golang-microservices/src/api/domain/github"
	"golang-microservices/src/api/domain/repositories"
	"golang-microservices/src/api/providers/github_provider"
	"golang-microservices/src/api/utils/errors"
	"strings"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	result := repositories.CreateRepoResponse{
		Id: response.Id,
		Name: response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
