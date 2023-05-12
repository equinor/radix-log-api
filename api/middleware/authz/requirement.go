package authz

import (
	apierrors "github.com/equinor/radix-log-api/api/errors"
)

// Requirement defines a requirement for use with policies.
type Requirement interface {
	// ValidateRequirement validates the requirement.
	// Returns an error if the requirement fails.
	ValidateRequirement(ctx *AuthorizationContext) error
}

// RequirementFunc wraps a function as a Requirement.
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
