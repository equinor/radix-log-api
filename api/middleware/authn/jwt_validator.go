package authn

import (
	"context"
	"net/url"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	apierrors "github.com/equinor/radix-log-api/api/errors"
)

type JwtValidator interface {
	Validate(token string) error
}

type jwtValidator struct {
	cacheProvider *jwks.CachingProvider
	issuerURL     string
	audience      string
}

func NewValidator(issuerUrl string, audience string) (JwtValidator, error) {
	url, err := url.Parse(issuerUrl)
	if err != nil {
		return nil, err
	}
	provider := jwks.NewCachingProvider(url, 5*time.Minute)
	return &jwtValidator{cacheProvider: provider, issuerURL: issuerUrl, audience: audience}, nil
}

func (v *jwtValidator) Validate(token string) error {
	validator, err := validator.New(
		v.cacheProvider.KeyFunc,
		validator.RS256,
		v.issuerURL,
		[]string{v.audience},
	)
	if err != nil {
		return err
	}
	_, err = validator.ValidateToken(context.TODO(), token)
	if err != nil {
		return apierrors.NewUnauthorizedError(apierrors.WithCause(err))
	}
	return nil
}
