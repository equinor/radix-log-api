package authz

import (
	"errors"
	"fmt"

	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/gin-gonic/gin"
)

var (
	errInvalidUserTypeInContext = errors.New("invalid user type")
	errPolicyNotFound           = func(policyName string) error { return fmt.Errorf("policy with name '%s' not found", policyName) }
)

var defaultPolicy = &policy{requirements: []Requirement{denyAnonymousUserRequirement}}

// Authorizer is used to create authorization middlewares using the Authorize() method.
type Authorizer interface {
	// Authorize returns a Gin middleware that validates the incoming request with the specified policies.
	// The middlesware aborts the Gin context if any of the policies fails validation.
	Authorize(policies ...string) gin.HandlerFunc
}

// AuthorizationConfiguration configures the authorizer.
type AuthorizationConfiguration interface {
	// AddPolicy registers a named policy.
	AddPolicy(name string, configure func(PolicyConfiguration)) AuthorizationConfiguration

	// WithDefaultPolicy sets the default policy.
	// This policy is used when no policy names are specified for the Authorize() method.
	WithDefaultPolicy(configure func(PolicyConfiguration)) AuthorizationConfiguration
}

// NewAuthorizer creates a new Authorizer
func NewAuthorizer(configure func(AuthorizationConfiguration)) Authorizer {
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
			if userTyped, ok := user.(authn.TokenPrincipal); ok {
				authCtx.user = userTyped
			} else {
				ctx.Error(errInvalidUserTypeInContext)
				ctx.Abort()
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
			if err := policy.ValidatePolicy(&authCtx); err != nil {
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
			return nil, errPolicyNotFound(policyName)
		}
		policies = append(policies, policy)
	}
	return policies, nil
}

func (a *authorizer) AddPolicy(name string, configure func(PolicyConfiguration)) AuthorizationConfiguration {
	p := &policy{}
	if configure != nil {
		configure(p)
	}
	a.policies[name] = p
	return a
}

func (a *authorizer) WithDefaultPolicy(configure func(PolicyConfiguration)) AuthorizationConfiguration {
	p := &policy{}
	if configure != nil {
		configure(p)
	}
	a.defaultPolicy = p
	return a
}
