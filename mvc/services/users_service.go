package services

import (
	"golang-microservices/mvc/domain"
	"golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	// Recupera os dados
	return domain.GetUser(userId)
}
