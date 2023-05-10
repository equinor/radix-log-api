package tests

import (
	"github.com/equinor/radix-log-api/tests/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	logService        *mock.MockLogService
	jwtValidator      *mock.MockJwtValidator
	applicationClient *mock.MockRadixApiApplicationClient
}

func (s *testSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.logService = mock.NewMockLogService(ctrl)
	s.jwtValidator = mock.NewMockJwtValidator(ctrl)
	s.applicationClient = mock.NewMockRadixApiApplicationClient(ctrl)
}
