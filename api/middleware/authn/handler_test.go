package authn

import (
	"errors"
	"testing"

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

func Test_Handler_EmptyProviderList_AnonymousUser(t *testing.T) {
	var ctx gin.Context
	sut := New()
	sut(&ctx)
	user, exist := ctx.Get(UserKey)
	require.True(t, exist)
	require.Implements(t, (*TokenPrincipal)(nil), user)
	assert.False(t, user.(TokenPrincipal).IsAuthenticated())
	assert.Empty(t, user.(TokenPrincipal).Token())
}

func Test_Handler_SingleProvider_UseReturnedToken(t *testing.T) {
	var ctx gin.Context
	tokenUser := &fakeTokenPrincipal{}
	ctrl := gomock.NewController(t)
	provider := NewMockAuthenticationProvider(ctrl)
	provider.EXPECT().Authenticate(ctx.Request).Return(tokenUser, nil).Times(1)
	sut := New(provider)
	sut(&ctx)
	user, exist := ctx.Get(UserKey)
	require.True(t, exist)
	assert.Equal(t, tokenUser, user)
}

func Test_Handler_SingleProvider_AnonymousUserWhenReturnNil(t *testing.T) {
	var ctx gin.Context
	ctrl := gomock.NewController(t)
	provider := NewMockAuthenticationProvider(ctrl)
	provider.EXPECT().Authenticate(ctx.Request).Return(nil, nil).Times(1)
	sut := New(provider)
	sut(&ctx)
	user, exist := ctx.Get(UserKey)
	require.True(t, exist)
	require.Implements(t, (*TokenPrincipal)(nil), user)
	assert.False(t, user.(TokenPrincipal).IsAuthenticated())
	assert.Empty(t, user.(TokenPrincipal).Token())
}

func Test_Handler_MultipleProvider_SkipProviderIterationOnFirstAuthenticatedUser(t *testing.T) {
	var ctx gin.Context
	ctrl := gomock.NewController(t)
	provider1 := NewMockAuthenticationProvider(ctrl)
	provider1.EXPECT().Authenticate(ctx.Request).Return(nil, nil).Times(1)
	provider2 := NewMockAuthenticationProvider(ctrl)
	provider2.EXPECT().Authenticate(ctx.Request).Return(&fakeTokenPrincipal{"", false}, nil).Times(1)
	provider3 := NewMockAuthenticationProvider(ctrl)
	provider3.EXPECT().Authenticate(ctx.Request).Return(&fakeTokenPrincipal{"token3", true}, nil).Times(1)
	provider4 := NewMockAuthenticationProvider(ctrl)
	sut := New(provider1, provider2, provider3, provider4)
	sut(&ctx)
	user, exist := ctx.Get(UserKey)
	require.True(t, exist)
	require.Implements(t, (*TokenPrincipal)(nil), user)
	assert.True(t, user.(TokenPrincipal).IsAuthenticated())
	assert.Equal(t, "token3", user.(TokenPrincipal).Token())
}

func Test_Handler_MultipleProvider_SkipProviderIterationWhenProviderReturnError(t *testing.T) {
	var ctx gin.Context
	providerErr := errors.New("err")
	ctrl := gomock.NewController(t)
	provider1 := NewMockAuthenticationProvider(ctrl)
	provider1.EXPECT().Authenticate(ctx.Request).Return(nil, nil).Times(1)
	provider2 := NewMockAuthenticationProvider(ctrl)
	provider2.EXPECT().Authenticate(ctx.Request).Return(nil, providerErr).Times(1)
	provider3 := NewMockAuthenticationProvider(ctrl)
	sut := New(provider1, provider2, provider3)
	sut(&ctx)
	assert.True(t, ctx.IsAborted())
	assert.Len(t, ctx.Errors, 1)
	assert.Equal(t, ctx.Errors[0].Err, providerErr)

}
