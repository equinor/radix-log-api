package authz

import (
	"net/http"
	"testing"

	"github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeTokenPrincipal struct {
	token           string
	isAuthenticated bool
}

func (t *fakeTokenPrincipal) IsAuthenticated() bool { return t.isAuthenticated }
func (t *fakeTokenPrincipal) Token() string         { return t.token }

func Test_Authorizer_DefaultConfiguration(t *testing.T) {
	authorizer := NewAuthorizer(func(ab AuthorizationConfiguration) {})
	t.Run("user context not set", func(t *testing.T) {
		t.Parallel()
		var ctx gin.Context
		authorizer.Authorize()(&ctx)
		assert.True(t, ctx.IsAborted())
		require.Len(t, ctx.Errors, 1)
		apiErr, ok := ctx.Errors[0].Err.(errors.APIStatus)
		require.True(t, ok)
		assert.Equal(t, http.StatusUnauthorized, apiErr.Status().Code)
	})
	t.Run("user context set to anonymous user", func(t *testing.T) {
		t.Parallel()
		var ctx gin.Context
		ctx.Set(authn.UserKey, &fakeTokenPrincipal{"", false})
		authorizer.Authorize()(&ctx)
		assert.True(t, ctx.IsAborted())
		require.Len(t, ctx.Errors, 1)
		apiErr, ok := ctx.Errors[0].Err.(errors.APIStatus)
		require.True(t, ok)
		assert.Equal(t, http.StatusUnauthorized, apiErr.Status().Code)
	})
	t.Run("user context set to invalid type", func(t *testing.T) {
		t.Parallel()
		var ctx gin.Context
		ctx.Set(authn.UserKey, "invalid data type")
		authorizer.Authorize()(&ctx)
		assert.True(t, ctx.IsAborted())
		require.Len(t, ctx.Errors, 1)
		assert.ErrorIs(t, ctx.Errors[0].Err, errInvalidUserTypeInContext)
	})
	t.Run("undefined policy name", func(t *testing.T) {
		t.Parallel()
		var ctx gin.Context
		authorizer.Authorize("unknown_policy")(&ctx)
		assert.True(t, ctx.IsAborted())
		require.Len(t, ctx.Errors, 1)
		expectedErr := errPolicyNotFound("unknown_policy")
		assert.ErrorContains(t, ctx.Errors[0].Err, expectedErr.Error())
	})
}
