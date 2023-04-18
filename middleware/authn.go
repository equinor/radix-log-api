package middleware

import (
	"github.com/equinor/radix-log-api/pkg/authn"
	"github.com/gin-gonic/gin"
)

type anonymousUser struct{}

func (u *anonymousUser) IsAuthenticated() bool  { return false }
func (u *anonymousUser) RawToken() string       { return "" }
func (u *anonymousUser) IsInRole(_ string) bool { return false }

func Authentication(auths ...authn.Provider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		setUserForContext(ctx, &anonymousUser{})
		for _, auth := range auths {
			user, err := auth.Authenticate(ctx.Request)
			if err != nil {
				ctx.AbortWithError(500, err)
				return
			}
			if user != nil {
				setUserForContext(ctx, user)
				// Skip other authenticators once an authenticated user is found in the request
				if user.IsAuthenticated() {
					return
				}
			}
		}
	}
}

func setUserForContext(ctx *gin.Context, user authn.ClaimsPrincipal) {
	ctx.Set("user", user)
}
