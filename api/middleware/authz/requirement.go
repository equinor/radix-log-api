package authz

import (
	apierrors "github.com/equinor/radix-log-api/api/errors"
)

type Requirement interface {
	ValidateRequirement(ctx *AuthorizationContext) error
}

type RequirementFunc func(ctx *AuthorizationContext) error

func (f RequirementFunc) ValidateRequirement(ctx *AuthorizationContext) error {
	return f(ctx)
}

var denyAnonymousUserRequirement RequirementFunc = func(ctx *AuthorizationContext) error {
	if user := ctx.User(); user == nil || !user.IsAuthenticated() {
		return apierrors.NewUnauthorizedError()
	}
	return nil
}
