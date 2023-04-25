package authz

import (
	apierrors "github.com/equinor/radix-log-api/api/errors"
)

type Requirement interface {
	Handle(ctx *AuthorizationContext) error
}

type RequirementFunc func(ctx *AuthorizationContext) error

func (f RequirementFunc) Handle(ctx *AuthorizationContext) error {
	return f(ctx)
}

type denyAnonymousUserRequirement struct{}

func (r *denyAnonymousUserRequirement) Handle(ctx *AuthorizationContext) error {
	if user := ctx.User(); user == nil || !user.IsAuthenticated() {
		return apierrors.NewUnauthorizedError()
	}
	return nil
}
