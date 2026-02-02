package authz

import (
	"errors"
	"net/http"
	"testing"

	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
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
		apiErr, ok := ctx.Errors[0].Err.(apierrors.APIStatus)
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
		apiErr, ok := ctx.Errors[0].Err.(apierrors.APIStatus)
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
	t.Run("user context set to authenticated user", func(t *testing.T) {
		t.Parallel()
		var ctx gin.Context
		ctx.Set(authn.UserKey, &fakeTokenPrincipal{"", true})
		authorizer.Authorize()(&ctx)
		assert.False(t, ctx.IsAborted())
		assert.Empty(t, ctx.Errors)
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

func Test_Authorizer_UseDefaultPolicy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defaultPolicy := NewMockPolicy(ctrl)
	defaultPolicy.EXPECT().ValidatePolicy(gomock.Any()).Return(nil).Times(1)
	authz := authorizer{defaultPolicy: defaultPolicy}
	var ctx gin.Context
	authz.Authorize()(&ctx)
}

func Test_Authorizer_UseNamedPolicyOnly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defaultPolicy := NewMockPolicy(ctrl)
	policy1 := NewMockPolicy(ctrl)
	policy1.EXPECT().ValidatePolicy(gomock.Any()).Return(nil).Times(1)
	policy2 := NewMockPolicy(ctrl)
	authz := authorizer{defaultPolicy: defaultPolicy, policies: map[string]Policy{"policy1": policy1, "policy2": policy2}}
	var ctx gin.Context
	authz.Authorize("policy1")(&ctx)
}

func Test_Authorizer_ValidateAllNamedPolicies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	policy1 := NewMockPolicy(ctrl)
	policy1.EXPECT().ValidatePolicy(gomock.Any()).Return(nil).Times(1)
	policy2 := NewMockPolicy(ctrl)
	policy2.EXPECT().ValidatePolicy(gomock.Any()).Return(nil).Times(1)
	authz := authorizer{policies: map[string]Policy{"policy1": policy1, "policy2": policy2}}
	var ctx gin.Context
	authz.Authorize("policy1", "policy2")(&ctx)
	assert.False(t, ctx.IsAborted())
	require.Empty(t, ctx.Errors)
}

func Test_Authorizer_SkipPolicyValidationAfterFirstError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	expectedErr := errors.New("first error")
	policy1 := NewMockPolicy(ctrl)
	policy1.EXPECT().ValidatePolicy(gomock.Any()).Return(expectedErr).Times(1)
	policy2 := NewMockPolicy(ctrl)
	authz := authorizer{policies: map[string]Policy{"policy1": policy1, "policy2": policy2}}
	var ctx gin.Context
	authz.Authorize("policy1", "policy2")(&ctx)
	assert.True(t, ctx.IsAborted())
	require.Len(t, ctx.Errors, 1)
	assert.ErrorIs(t, ctx.Errors[0].Err, expectedErr)
}
