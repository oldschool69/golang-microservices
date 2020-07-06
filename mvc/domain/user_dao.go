package domain

import (
	"errors"
	"fmt"
)

// Camada de domain, única camada responsável pelo acesso aos dados
// como consulta a dados etc

// Dados mockados para teste
var (
	users = map[int64]*User{
		123: {
			Id: 123,
			FirstName: "Flavio",
			LastName: "Oliveira",
			Email: "flavio.marcondes@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, error) {
	if 	user := users[userId]; user  != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("user %v was not found\n", userId))
}
