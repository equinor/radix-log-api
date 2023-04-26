package authz

import (
	apierrors "github.com/equinor/radix-log-api/api/errors"
)

type Requirement interface {
	HandleRequirement(ctx *AuthorizationContext) error
}

type RequirementFunc func(ctx *AuthorizationContext) error

func (f RequirementFunc) HandleRequirement(ctx *AuthorizationContext) error {
	return f(ctx)
}

var denyAnonymousUserRequirement RequirementFunc = func(ctx *AuthorizationContext) error {
	if user := ctx.User(); user == nil || !user.IsAuthenticated() {
		return apierrors.NewUnauthorizedError()
	}
	return nil
}
