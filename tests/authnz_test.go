package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/router"
	"github.com/equinor/radix-log-api/tests/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func Test_AuthnzTestSuite(t *testing.T) {
	suite.Run(t, new(authnzTestSuite))
}

type authnzTestSuite struct {
	suite.Suite
	logService        *mock.MockLogService
	jwtValidator      *mock.MockJwtValidator
	applicationClient *mock.MockRadixApiApplicationClient
}

func (s *authnzTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.logService = mock.NewMockLogService(ctrl)
	s.jwtValidator = mock.NewMockJwtValidator(ctrl)
	s.applicationClient = mock.NewMockRadixApiApplicationClient(ctrl)
}

func (s *authnzTestSuite) Test_MissingAuthorizationHeader() {
	r, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	req, _ := newInventoryRequest("anyapp", "anyenv", "anycomp")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnzTestSuite) Test_MissingBearerInAuthorizationHeader() {
	r, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	req, _ := newInventoryRequest("anyapp", "anyenv", "anycomp", withAuthorization(""))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnzTestSuite) Test_MissingTokenInBearerAuthorizationHeader() {
	r, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	req, _ := newInventoryRequest("anyapp", "anyenv", "anycomp", withBearerAuthorization(""))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnzTestSuite) Test_JwtValidatorTokenUnauthorized() {
	r, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	token := "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(apierrors.NewUnauthorizedError()).Times(1)
	req, _ := newInventoryRequest("anyapp", "anyenv", "anycomp", withBearerAuthorization(token))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnzTestSuite) Test_JwtValidatorTokenGenericError() {
	r, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	token := "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(errors.New("generic error")).Times(1)
	req, _ := newInventoryRequest("anyapp", "anyenv", "anycomp", withBearerAuthorization(token))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *authnzTestSuite) Test_SuccessfulAuthentication() {
	r, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	token := "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.applicationClient.EXPECT().GetApplication(&getApplicationMatcher{}, &getApplicationAuthMatcher{}).Return(nil, nil).Times(1)
	req, _ := newInventoryRequest("anyapp", "anyenv", "anycomp", withBearerAuthorization(token))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

type getApplicationMatcher struct{}

// Matches returns whether x is a match.
func (m *getApplicationMatcher) Matches(x interface{}) bool {
	return true
}

// String describes what the matcher matches.
func (m *getApplicationMatcher) String() string {
	return ""
}

type getApplicationAuthMatcher struct{}

// Matches returns whether x is a match.
func (m *getApplicationAuthMatcher) Matches(x interface{}) bool {
	return true
}

// String describes what the matcher matches.
func (m *getApplicationAuthMatcher) String() string {
	return ""
}
