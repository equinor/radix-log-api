package router

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/equinor/radix-log-api/api/controllers"
	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/equinor/radix-log-api/internal/tests/mock"
	"github.com/equinor/radix-log-api/internal/tests/request"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func Test_AuthnTestSuite(t *testing.T) {
	suite.Run(t, new(authnTestSuite))
}

type authnTestSuite struct {
	suite.Suite
	JwtValidator      *authn.MockJwtValidator
	ApplicationClient *mock.MockRadixApiApplicationClient
	ApiController     *controllers.MockController
}

func (s *authnTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.JwtValidator = authn.NewMockJwtValidator(ctrl)
	s.ApplicationClient = mock.NewMockRadixApiApplicationClient(ctrl)
	s.ApiController = controllers.NewMockController(ctrl)
}

func (s *authnTestSuite) sut() http.Handler {
	s.ApiController.EXPECT().Endpoints().Return([]controllers.Endpoint{{Path: "any", Method: "GET", Handler: func(ctx *gin.Context) { ctx.String(200, "apiresponse") }}}).Times(1)
	sut, err := New(s.JwtValidator, s.ApplicationClient, s.ApiController)
	s.Require().NoError(err)
	return sut
}

func (s *authnTestSuite) Test_MissingAuthorizationHeader() {
	req, _ := request.New("/api/v1/any")
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnTestSuite) Test_MissingBearerInAuthorizationHeader() {
	req, _ := request.New("/api/v1/any", request.WithAuthorization(""))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnTestSuite) Test_MissingTokenInBearerAuthorizationHeader() {
	req, _ := request.New("/api/v1/any", request.WithBearerAuthorization(""))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnTestSuite) Test_JwtValidatorTokenUnauthorized() {
	token := "anytoken"
	s.JwtValidator.EXPECT().Validate(gomock.Any(), token).Return(apierrors.NewUnauthorizedError()).Times(1)

	req, _ := request.New("/api/v1/any", request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authnTestSuite) Test_JwtValidatorTokenGenericError() {
	token := "anytoken"
	s.JwtValidator.EXPECT().Validate(gomock.Any(), token).Return(errors.New("generic error")).Times(1)

	req, _ := request.New("/api/v1/any", request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *authnTestSuite) Test_SuccessfulAuthentication() {
	token := "anytoken"
	s.JwtValidator.EXPECT().Validate(gomock.Any(), token).Return(nil).Times(1)

	req, _ := request.New("/api/v1/any", request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal("apiresponse", string(actual))
}
