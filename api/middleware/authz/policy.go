package authz

import (
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/gin-gonic/gin"
)

// AuthorizationContext used as arguments to policies and requirements
type AuthorizationContext struct {
	user   authn.TokenPrincipal
	ginCtx *gin.Context
}

// User returns the user for the request
func (ctx *AuthorizationContext) User() authn.TokenPrincipal {
	return ctx.user
}

// GinCtx returnes the Gin context for the request
func (ctx *AuthorizationContext) GinCtx() *gin.Context {
	return ctx.ginCtx
}

// Policy defines a set of requirements that must be fulfilled in order to process the request.
type Policy interface {
	// ValidatePolicy validates the requirements defined for the policy.
	// Returns an error if one or more requirements fails.
	ValidatePolicy(ctx *AuthorizationContext) error
}

// PolicyConfiguration configures a policy.
type PolicyConfiguration interface {
	// RequireAuthenticatedUser adds a requirement to require an authenticated user in the request.
	RequireAuthenticatedUser() PolicyConfiguration

	// AddRequirement adds a requirement to the policy
	AddRequirement(Requirement) PolicyConfiguration
}

type policy struct {
	requirements []Requirement
}

func (p *policy) ValidatePolicy(ctx *AuthorizationContext) error {
	for _, r := range p.requirements {
		if err := r.ValidateRequirement(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (p *policy) RequireAuthenticatedUser() PolicyConfiguration {
	p.requirements = append(p.requirements, denyAnonymousUserRequirement)
	return p
}

func (p *policy) AddRequirement(r Requirement) PolicyConfiguration {
	p.requirements = append(p.requirements, r)
	return p
}
