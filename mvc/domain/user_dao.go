package domain

import (
	"fmt"
	"golang-microservices/mvc/utils"
	"net/http"
)

// Camada de domain, única camada responsável pelo acesso aos dados
// como consulta a dados etc

// Dados mockados para teste
var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Flavio",
			LastName:  "Oliveira",
			Email:     "flavio.marcondes@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}
}
