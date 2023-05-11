package logs

import (
	"net/http"

	"github.com/equinor/radix-log-api/api/router"
	"github.com/equinor/radix-log-api/internal/tests/mock"
	"github.com/equinor/radix-log-api/internal/tests/request"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type controllerTestSuite struct {
	suite.Suite
	LogService        *mock.MockLogService
	JwtValidator      *mock.MockJwtValidator
	ApplicationClient *mock.MockRadixApiApplicationClient
}

func (s *controllerTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.LogService = mock.NewMockLogService(ctrl)
	s.JwtValidator = mock.NewMockJwtValidator(ctrl)
	s.ApplicationClient = mock.NewMockRadixApiApplicationClient(ctrl)
}

func (s *controllerTestSuite) sut() http.Handler {
	sut, err := router.New(s.JwtValidator, s.ApplicationClient, New(s.LogService))
	s.Require().NoError(err)
	return sut
}

func (s *controllerTestSuite) newRequest(url string) *http.Request {
	req, err := request.New(url, request.WithBearerAuthorization("anytoken"))
	s.Require().NoError(err)
	return req
}
