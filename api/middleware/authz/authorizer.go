package authz

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/gin-gonic/gin"
)

type Authorizer interface {
	Authorize(policies ...string) gin.HandlerFunc
}

type AuthorizationBuilder interface {
	AddPolicy(name string, configure func(PolicyBuilder)) AuthorizationBuilder
	WithDefaultPolicy(configure func(PolicyBuilder)) AuthorizationBuilder
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
		user, exists := ctx.Get(authn.UserKey)
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
			if err := policy.HandlePolicy(&authCtx); err != nil {
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
