package oauth

import (
	"github.com/gin-gonic/gin"
	"golang-microservices/oauth-api/src/api/domain/oauth"
	"golang-microservices/oauth-api/src/api/services"
	"golang-microservices/src/api/utils/errors"
	"net/http"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccesTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
	}
	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context) {
	token, err := services.OauthService.GetAccessToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}
