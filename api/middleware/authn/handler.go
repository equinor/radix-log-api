package authn

import (
	"github.com/gin-gonic/gin"
)

const (
	UserKey = "user"
)

var _ TokenPrincipal = &anonymousPrincipal{}

type anonymousPrincipal struct{}

func (u *anonymousPrincipal) IsAuthenticated() bool { return false }
func (u *anonymousPrincipal) Token() string         { return "" }

func New(auths ...AuthenticationProvider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		setUserForContext(ctx, &anonymousPrincipal{})
		for _, auth := range auths {
			user, err := auth.Authenticate(ctx.Request)
			if err != nil {
				_ = ctx.Error(err)
				ctx.Abort()
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

func setUserForContext(ctx *gin.Context, user TokenPrincipal) {
	ctx.Set(UserKey, user)
}
