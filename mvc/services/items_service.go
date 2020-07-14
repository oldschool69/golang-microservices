package services

import (
	"golang-microservices/mvc/domain"
	"golang-microservices/mvc/utils"
	"net/http"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (i *itemsService) GetItems(itemId string) (*domain.Item, *utils.ApplicationError) {
	// Recupera os dados
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
