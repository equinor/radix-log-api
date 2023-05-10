package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/router"
	"github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	"github.com/equinor/radix-log-api/tests/internal/match"
	"github.com/stretchr/testify/suite"
)

func Test_AuthnzTest(t *testing.T) {
	suite.Run(t, new(authnzTestSuite))
}

type authnzTestSuite struct {
	testSuite
}

func (s *authnzTestSuite) Test_MissingAuthorizationHeader() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	req, _ := newRequest(newComponentInventoryUrl("anyapp", "anyenv", "anycomp"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnzTestSuite) Test_MissingBearerInAuthorizationHeader() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	req, _ := newRequest(newComponentInventoryUrl("anyapp", "anyenv", "anycomp"), withAuthorization(""))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnzTestSuite) Test_MissingTokenInBearerAuthorizationHeader() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	req, _ := newRequest(newComponentInventoryUrl("anyapp", "anyenv", "anycomp"), withBearerAuthorization(""))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnzTestSuite) Test_JwtValidatorTokenUnauthorized() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	token := "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(apierrors.NewUnauthorizedError()).Times(1)

	req, _ := newRequest(newComponentInventoryUrl("anyapp", "anyenv", "anycomp"), withBearerAuthorization(token))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnzTestSuite) Test_JwtValidatorTokenGenericError() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	token := "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(errors.New("generic error")).Times(1)

	req, _ := newRequest(newComponentInventoryUrl("anyapp", "anyenv", "anycomp"), withBearerAuthorization(token))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *authnzTestSuite) Test_Authorization_GetApplication_AppNotFound() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, token := "anyapp", "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.applicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, &application.GetApplicationNotFound{}).Times(1)

	req, _ := newRequest(newComponentInventoryUrl(appName, "anyenv", "anycomp"), withBearerAuthorization(token))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusForbidden, w.Code)
}

func (s *authnzTestSuite) Test_Authorization_GetApplication_Unauthorized() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, token := "anyapp", "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.applicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, &application.GetApplicationUnauthorized{}).Times(1)

	req, _ := newRequest(newComponentInventoryUrl(appName, "anyenv", "anycomp"), withBearerAuthorization(token))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusForbidden, w.Code)
}

func (s *authnzTestSuite) Test_Authorization_GetApplication_Forbidden() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, token := "anyapp", "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.applicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, &application.GetApplicationForbidden{}).Times(1)

	req, _ := newRequest(newComponentInventoryUrl(appName, "anyenv", "anycomp"), withBearerAuthorization(token))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusForbidden, w.Code)
}

func (s *authnzTestSuite) Test_Authorization_GetApplication_InternalServerError() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, token := "anyapp", "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.applicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, &application.GetApplicationInternalServerError{}).Times(1)

	req, _ := newRequest(newComponentInventoryUrl(appName, "anyenv", "anycomp"), withBearerAuthorization(token))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *authnzTestSuite) Test_Authorization_GetApplication_Conflict() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, token := "anyapp", "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.applicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, &application.GetApplicationConflict{}).Times(1)

	req, _ := newRequest(newComponentInventoryUrl(appName, "anyenv", "anycomp"), withBearerAuthorization(token))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *authnzTestSuite) Test_Authorization_GetApplication_GenericError() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, token := "anyapp", "anytoken"
	s.jwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.applicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, errors.New("any error")).Times(1)

	req, _ := newRequest(newComponentInventoryUrl(appName, "anyenv", "anycomp"), withBearerAuthorization(token))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}
