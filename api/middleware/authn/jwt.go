package authn

import (
	"net/http"
	"strings"

	"github.com/equinor/radix-log-api/pkg/jwt"
)

func NewJwt(validator jwt.Validator) AuthenticationProvider {
	return &jwtProvider{validator: validator}
}

type jwtPrincipal struct {
	token string
}

func (p *jwtPrincipal) Token() string         { return p.token }
func (p *jwtPrincipal) IsAuthenticated() bool { return true }

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
		return nil, nil
	}
	token := strings.TrimSpace(authParts[1])
	if len(token) == 0 {
		return nil, nil
	}
	valid, err := a.validator.Validate(token)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, nil
	}
	return &jwtPrincipal{token: token}, nil
}
