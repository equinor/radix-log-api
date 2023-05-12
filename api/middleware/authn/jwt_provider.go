package authn

import (
	"net/http"
	"strings"
)

type jwtPrincipal struct {
	token string
}

func (p *jwtPrincipal) Token() string         { return p.token }
func (p *jwtPrincipal) IsAuthenticated() bool { return true }

func NewJwtProvider(validator JwtValidator) AuthenticationProvider {
	return &jwtProvider{validator: validator}
}

type jwtProvider struct {
	validator JwtValidator
}

func (a *jwtProvider) Authenticate(req *http.Request) (TokenPrincipal, error) {
	authorization := req.Header.Get("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		return nil, nil
	}
	authParts := strings.Split(authorization, "Bearer ")
	if len(authParts) != 2 {
		return nil, nil
	}
	token := strings.TrimSpace(authParts[1])
	if len(token) == 0 {
		return nil, nil
	}
	err := a.validator.Validate(token)
	if err != nil {
		return nil, err
	}
	return &jwtPrincipal{token: token}, nil
}
