package services

import "golang-microservices/mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	// Recupera os dados
	return domain.GetUser(userId)
}
