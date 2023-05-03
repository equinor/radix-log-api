package authz

import (
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/gin-gonic/gin"
)

var defaultPolicy = &policy{requirements: []Requirement{denyAnonymousUserRequirement}}

type AuthorizationContext struct {
	user   authn.TokenPrincipal
	ginCtx *gin.Context
}

func (ctx *AuthorizationContext) User() authn.TokenPrincipal {
	return ctx.user
}

func (ctx *AuthorizationContext) GinCtx() *gin.Context {
	return ctx.ginCtx
}

type Policy interface {
	HandlePolicy(ctx *AuthorizationContext) error
}

type PolicyBuilder interface {
	RequireAuthenticatedUser() PolicyBuilder
	AddRequirement(Requirement) PolicyBuilder
}

type policy struct {
	requirements []Requirement
}

func (p *policy) HandlePolicy(ctx *AuthorizationContext) error {
	for _, r := range p.requirements {
		if err := r.HandleRequirement(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (p *policy) RequireAuthenticatedUser() PolicyBuilder {
	p.requirements = append(p.requirements, denyAnonymousUserRequirement)
	return p
}

func (p *policy) AddRequirement(r Requirement) PolicyBuilder {
	p.requirements = append(p.requirements, r)
	return p
}
