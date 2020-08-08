package oauth

import (
	"fmt"
	"golang-microservices/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.ApiError {
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserId)
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.ApiError) {
	token := tokens[accessToken]
	if token == nil {
		return nil, errors.NewNotFoundApiError("no access token found with given parameter")
	}
	return token, nil
}
