package controllers

import (
	"encoding/json"
	"golang-microservices/mvc/services"
	"net/http"
	"strconv"
)

// Única responsabilidade do controller é validar request e enviar resposta para os clientes
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		// return Bad quest to the client
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user_id must be a number\n"))
		return
	}

	// Consulta a camada de serviço para pegar os dados
	user, err := services.GetUser(userId)
	if err != nil {
		// Handle the error and return to the client
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}