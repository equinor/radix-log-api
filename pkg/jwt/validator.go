package jwt

import (
	"context"
	"net/url"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type Validator interface {
	Validate(token string) (bool, error)
}

type jwtValidator struct {
	cacheProvider *jwks.CachingProvider
	issuerURL     string
	audience      string
}

func NewValidator(issuerUrl string, audience string) (Validator, error) {
	url, err := url.Parse(issuerUrl)
	if err != nil {
		return nil, err
	}
	provider := jwks.NewCachingProvider(url, 5*time.Minute)
	return &jwtValidator{cacheProvider: provider, issuerURL: issuerUrl, audience: audience}, nil
}

func (v *jwtValidator) Validate(token string) (bool, error) {
	validator, err := validator.New(
		v.cacheProvider.KeyFunc,
		validator.RS256,
		v.issuerURL,
		[]string{v.audience},
	)
	if err != nil {
		return false, err
	}
	_, err = validator.ValidateToken(context.TODO(), token)
	return err == nil, nil
}
