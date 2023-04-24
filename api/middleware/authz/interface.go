package authz

import (
	"errors"
	"fmt"
	"net/http"

	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/middleware/authn"

	"github.com/gin-gonic/gin"
)

var defaultPolicy = &policy{requirements: []Requirement{&denyAnonymousUserRequirement{}}}

type AuthorizationContext struct {
	user   authn.ClaimsPrincipal
	ginCtx *gin.Context
}

func (ctx *AuthorizationContext) User() authn.ClaimsPrincipal {
	return ctx.user
}

func (ctx *AuthorizationContext) GinCtx() *gin.Context {
	return ctx.ginCtx
}

type Authorizer interface {
	Authorize(policies ...string) gin.HandlerFunc
}

type AuthorizationBuilder interface {
	AddPolicy(name string, configure func(PolicyBuilder)) AuthorizationBuilder
	WithDefaultPolicy(configure func(PolicyBuilder)) AuthorizationBuilder
}

type Policy interface {
	Handle(ctx *AuthorizationContext) error
}

type PolicyBuilder interface {
	RequireAuthenticatedUser() PolicyBuilder
	AddRequirement(Requirement) PolicyBuilder
}

type RequirementFunc func(ctx *AuthorizationContext) error

func (f RequirementFunc) Handle(ctx *AuthorizationContext) error {
	return f(ctx)
}

type Requirement interface {
	Handle(ctx *AuthorizationContext) error
}

func NewAuthorizer(configure func(AuthorizationBuilder)) Authorizer {
	authz := &authorizer{
		policies:      map[string]Policy{},
		defaultPolicy: defaultPolicy,
	}

	if configure != nil {
		configure(authz)
	}

	return authz
}

type authorizer struct {
	policies      map[string]Policy
	defaultPolicy Policy
}

func (a *authorizer) Authorize(policyNames ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authCtx := AuthorizationContext{ginCtx: ctx}
		user, exists := ctx.Get("user")
		if exists {
			if userTyped, ok := user.(authn.ClaimsPrincipal); ok {
				authCtx.user = userTyped
			} else {
				ctx.AbortWithError(http.StatusInternalServerError, errors.New("invalid user type"))
				return
			}
		}

		policies, err := a.getPoliciesByName(policyNames)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		if len(policies) == 0 && a.defaultPolicy != nil {
			policies = append(policies, a.defaultPolicy)
		}

		for _, policy := range policies {
			if err := policy.Handle(&authCtx); err != nil {
				ctx.Error(err)
				ctx.Abort()
				return
			}
		}
	}
}

func (a *authorizer) getPoliciesByName(policyNames []string) ([]Policy, error) {
	var policies []Policy

	for _, policyName := range policyNames {
		policy, found := a.policies[policyName]
		if !found {
			return nil, fmt.Errorf("policy with name '%s' not found", policyName)
		}
		policies = append(policies, policy)
	}
	return policies, nil
}

func (a *authorizer) AddPolicy(name string, configure func(PolicyBuilder)) AuthorizationBuilder {
	p := &policy{}
	if configure != nil {
		configure(p)
	}
	a.policies[name] = p
	return a
}

func (a *authorizer) WithDefaultPolicy(configure func(PolicyBuilder)) AuthorizationBuilder {
	p := &policy{}
	if configure != nil {
		configure(p)
	}
	a.defaultPolicy = p
	return a
}

type policy struct {
	requirements []Requirement
}

func (p *policy) Handle(ctx *AuthorizationContext) error {
	for _, r := range p.requirements {
		if err := r.Handle(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (p *policy) RequireAuthenticatedUser() PolicyBuilder {
	p.requirements = append(p.requirements, &denyAnonymousUserRequirement{})
	return p
}

func (p *policy) AddRequirement(r Requirement) PolicyBuilder {
	p.requirements = append(p.requirements, r)
	return p
}

type denyAnonymousUserRequirement struct{}

func (r *denyAnonymousUserRequirement) Handle(ctx *AuthorizationContext) error {
	if user := ctx.User(); user == nil || !user.IsAuthenticated() {
		return apierrors.NewUnauthorizedError()
	}
	return nil
}
