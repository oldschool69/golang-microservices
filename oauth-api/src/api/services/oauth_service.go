package services

import (
	"golang-microservices/oauth-api/src/api/domain/oauth"
	"golang-microservices/src/api/utils/errors"
	"time"
)

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccesTokenRequest) (*oauth.AccessToken, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccesTokenRequest) (*oauth.AccessToken, errors.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, err := oauth.GetUserByUserNameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	token := oauth.AccessToken{
		UserId:  user.Id,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}
	if err := token.Save(); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)
	if err != nil {
		return nil, err
	}
	if token.IsExpired() {
		return nil, errors.NewNotFoundApiError("no access token with given parameter")
	}
	return token, nil
}
