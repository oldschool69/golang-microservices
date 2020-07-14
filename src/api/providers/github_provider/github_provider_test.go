package github_provider

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang-microservices/src/api/clients/restclient"
	"golang-microservices/src/api/domain/github"
	"net/http"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient response", err.Message)

}

func TestCreateRepoErrorInvalidResponseBody(t *testing.T) {
	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient response", err.Message)

}
