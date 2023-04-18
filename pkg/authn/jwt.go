package authn

import (
	"errors"
	"net/http"
	"strings"

	"github.com/equinor/radix-log-api/pkg/jwt"
)

func NewJwt(validator jwt.Validator) Provider {
	return &jwtProvider{validator: validator}
}

type jwtProvider struct {
	validator jwt.Validator
}

func (a *jwtProvider) Authenticate(req *http.Request) (ClaimsPrincipal, error) {
	authorization := req.Header.Get("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		return nil, nil
	}
	authParts := strings.Split(authorization, "Bearer ")
	if len(authParts) != 2 {
		return nil, errors.New("invalid Authorization header")
	}
	token := strings.TrimSpace(authParts[1])
	if len(token) == 0 {
		return nil, errors.New("invalid Authorization header")
	}
	err := a.validator.Validate(token)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
