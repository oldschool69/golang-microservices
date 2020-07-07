package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, status int, body interface{}) {
	// Dependendo to tipo do header do request, retorna
	// formato XML ou JSON
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

func RespondError(c *gin.Context, err *ApplicationError) {
	// Dependendo to tipo do header do request, retorna
	// formato XML ou JSON
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
		return
	}
	c.JSON(err.StatusCode, err)
}
