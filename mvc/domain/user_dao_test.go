package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization

	// Execution
	user, err := UserDao.GetUser(0)

	//Validation
	assert.Nil(t, user, "User with Id 0 not expected")
	assert.NotNil(t, err, "Error should not be nil for User Id 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Flavio", user.FirstName)
	assert.EqualValues(t, "Oliveira", user.LastName)
	assert.EqualValues(t, "flavio.marcondes@gmail.com", user.Email)

}
