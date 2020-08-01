package main

import (
	"bufio"
	"fmt"
	"golang-microservices/src/api/domain/repositories"
	"golang-microservices/src/api/services"
	"golang-microservices/src/api/utils/errors"
	"os"
	"sync"
)

var (
	success map[string]string
	fail    map[string]errors.ApiError
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result *repositories.CreateRepoResponse
	Error  errors.ApiError
}

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("/home/flavio/go/src/golang-microservices/concurrency/requests.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name: line,
		}
		result = append(result, request)
	}
	return result
}

func main() {
	requests := getRequests()

	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))
	success = make(map[string]string)
	fail    = make(map[string]errors.ApiError)
	input := make(chan createRepoResult)
	// Esse técnica utiliza um buffered channel para
	// limitar a quantidade de go routines que são
	// criadas, nesse caso 10 por vez
	buffer := make(chan bool, 10)
	var wg sync.WaitGroup

	go handleResults(&wg, input)
	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(buffer, input, request)
	}

	// Nessse ponto é necessário esperar cada resultado ser processado pela
	// função handleResults antes de fechar o channel
	wg.Wait()
	close(input)

	for k, v := range success {
		fmt.Println("for reponame ", k)
		fmt.Println("success result was ", v)
	}

	for k, v := range fail {
		fmt.Println("for reponame ", k)
		fmt.Println("fail result was ", v)
	}


}

func handleResults(wg *sync.WaitGroup, input chan createRepoResult) {
	for result := range input {
		if result.Error != nil {
			fail[result.Request.Name] = result.Error
	 	} else {
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done()
	}
}

func createRepo(buffer chan bool, output chan createRepoResult, request repositories.CreateRepoRequest) {
	result, err := services.RepositoryService.CreateRepo(request)
	repoResult := createRepoResult{
		Request: request,
		Result: result,
		Error:  err,
	}
	output <- repoResult
	<-buffer
}
