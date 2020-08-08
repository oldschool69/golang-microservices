package oauth

import "golang-microservices/src/api/utils/errors"

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username, FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"oldschool69": {Id: 123, Username: "oldschool69"},
	}
)

func GetUserByUserNameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundApiError("No user found with given parameters")
	}
	return user, nil
}
