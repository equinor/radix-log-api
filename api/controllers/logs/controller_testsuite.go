package logs

import (
	"net/http"

	"github.com/equinor/radix-common/utils/pointers"
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/equinor/radix-log-api/api/router"
	"github.com/equinor/radix-log-api/internal/tests/request"
	"github.com/equinor/radix-log-api/pkg/authz/requirement"
	"github.com/equinor/radix-log-api/pkg/radixapi/models"
	logsservice "github.com/equinor/radix-log-api/pkg/services/logs"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

var anyApp = &models.Application{Name: pointers.Ptr("anyapp"), Registration: &models.ApplicationRegistration{AppID: pointers.Ptr("some-random-id")}}

type controllerTestSuite struct {
	suite.Suite
	LogService   *logsservice.MockLogService
	JwtValidator *authn.MockJwtValidator
	AppProvider  *requirement.MockRadixAppProvider
}

func (s *controllerTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.LogService = logsservice.NewMockLogService(ctrl)
	s.JwtValidator = authn.NewMockJwtValidator(ctrl)
	s.AppProvider = requirement.NewMockRadixAppProvider(ctrl)
}

func (s *controllerTestSuite) sut() http.Handler {
	sut, err := router.New(s.JwtValidator, s.AppProvider, New(s.LogService))
	s.Require().NoError(err)
	return sut
}

func (s *controllerTestSuite) newRequest(url string) *http.Request {
	req, err := request.New(url, request.WithBearerAuthorization("anytoken"))
	s.Require().NoError(err)
	return req
}
