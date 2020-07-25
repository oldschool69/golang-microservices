package repositories

import (
	"github.com/gin-gonic/gin"
	"golang-microservices/src/api/domain/repositories"
	"golang-microservices/src/api/services"
	"golang-microservices/src/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}
	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}
	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(result.StatusCode, result)
}

