package services

import (
	"github.com/stretchr/testify/assert"
	"golang-microservices/mvc/domain"
	"golang-microservices/mvc/utils"
	"net/http"
	"testing"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

// Força chamar a função GetUser da interface ao invés de chamar
// a implementação em user_dao.go
// Fazendo dessa forma, ao invés de executar uma consulta a uma
// fonte de dados real, o que não é indicado em testes unitários,
// mocka-se os dados executando a implementação local da função
// GetUser implementada através da interface domain.usersDaoInterface
func init() {
	//seta a interface na camada de domain para usar a função local
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 was not found",
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}
	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
