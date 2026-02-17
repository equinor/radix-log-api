package router

import (
	"testing"

	"github.com/equinor/radix-log-api/api/controllers"
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/equinor/radix-log-api/pkg/authz/requirement"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_RouterMapsControllers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiController := controllers.NewMockController(ctrl)
	apiController.EXPECT().Endpoints().Times(1)
	_, err := New(authn.NewMockJwtValidator(ctrl), requirement.NewMockRadixAppProvider(ctrl), apiController)
	assert.NoError(t, err)
}
