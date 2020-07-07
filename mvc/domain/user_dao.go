package domain

import (
	"fmt"
	"golang-microservices/mvc/utils"
	"log"
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

	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}
type userDao struct {}


func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("*** We're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
