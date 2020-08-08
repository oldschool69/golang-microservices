package services

import (
	"golang-microservices/src/api/config"
	"golang-microservices/src/api/domain/github"
	"golang-microservices/src/api/domain/repositories"
	"golang-microservices/src/api/providers/github_provider"
	"golang-microservices/src/api/utils/errors"
	"net/http"
	"sync"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {

	if err := input.Validate(); err != nil {
		return nil, err
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
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *repoService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	inputChannel := make(chan repositories.CreateRepositoriesResult)
	outputChannel := make(chan repositories.CreateReposResponse)
	defer close(outputChannel)
	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, inputChannel, outputChannel)

	for _, result := range requests {
		wg.Add(1)
		go s.CreateRepoConcurrent(result, inputChannel)
	}

	wg.Wait()
	close(inputChannel)

	responses := <-outputChannel

	successfullyRequests := 0
	for _, response := range responses.Results {
		if response.Response != nil {
			successfullyRequests++
		}
	}
	if successfullyRequests == 0 {
		responses.StatusCode = responses.Results[0].Error.GetStatus()
	} else if successfullyRequests == len(requests) {
		responses.StatusCode = http.StatusCreated
	} else {
		responses.StatusCode = http.StatusPartialContent
	}
	return responses, nil
}

func (s *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var responses repositories.CreateReposResponse

	for result := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: result.Response,
			Error:    result.Error,
		}
		responses.Results = append(responses.Results, repoResult)
		wg.Done()
	}

	output <- responses
}

func (s *repoService) CreateRepoConcurrent(request repositories.CreateRepoRequest, input chan repositories.CreateRepositoriesResult) {
	if err := request.Validate(); err != nil {
		input <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo(request)

	if err != nil {
		input <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	input <- repositories.CreateRepositoriesResult{Response: result}
}
