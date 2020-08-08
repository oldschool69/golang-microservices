package oauth

import (
	"golang-microservices/src/api/utils/errors"
	"strings"
)

type AccesTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *AccesTokenRequest) Validate() errors.ApiError {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return errors.NewBadRequestError("invalid username")
	}

	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
