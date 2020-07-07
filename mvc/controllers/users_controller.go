package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-microservices/mvc/services"
	"golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

// Única responsabilidade do controller é validar request e enviar resposta para os clientes
func GetUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		// return Bad quest to the client
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_Request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	// Consulta a camada de serviço para pegar os dados
	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		// Handle the error and return to the client
		utils.RespondError(c, apiErr)
		return
	}

	// return user to client
	utils.Respond(c, http.StatusOK, user)
}
