package router

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/equinor/radix-log-api/api/controllers"
	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/middleware/authn"
	"github.com/equinor/radix-log-api/internal/tests/match"
	"github.com/equinor/radix-log-api/internal/tests/mock"
	"github.com/equinor/radix-log-api/internal/tests/request"
	"github.com/equinor/radix-log-api/pkg/constants"
	"github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

var testAuthzEndpoints []controllers.Endpoint = []controllers.Endpoint{
	{
		Path:                  "appadminpolicy",
		Method:                "GET",
		Handler:               func(ctx *gin.Context) { ctx.Status(200) },
		AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
	},
	{
		Path:                  "appadminpolicy/:appName",
		Method:                "GET",
		Handler:               func(ctx *gin.Context) { ctx.Status(200) },
		AuthorizationPolicies: []string{constants.AuthorizationPolicyAppAdmin},
	},
	{
		Path:                  "invalidpolicy",
		Method:                "GET",
		Handler:               func(ctx *gin.Context) { ctx.Status(200) },
		AuthorizationPolicies: []string{"non-existing-policy"},
	},
	{
		Path:    "defaultpolicy",
		Method:  "GET",
		Handler: func(ctx *gin.Context) { ctx.Status(200) },
	},
}

func Test_AuthzTestSuite(t *testing.T) {
	suite.Run(t, new(authzTestSuite))
}

type authzTestSuite struct {
	suite.Suite
	JwtValidator      *authn.MockJwtValidator
	ApplicationClient *mock.MockRadixApiApplicationClient
	ApiController     *mock.MockController
}

func (s *authzTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.JwtValidator = authn.NewMockJwtValidator(ctrl)
	s.ApplicationClient = mock.NewMockRadixApiApplicationClient(ctrl)
	s.ApiController = mock.NewMockController(ctrl)
}

func (s *authzTestSuite) sut() http.Handler {
	s.ApiController.EXPECT().Endpoints().Return(testAuthzEndpoints).Times(1)
	sut, err := New(s.JwtValidator, s.ApplicationClient, s.ApiController)
	s.Require().NoError(err)
	return sut
}

func (s *authzTestSuite) Test_Authorization_AddAdminPolicy_AppNotfound() {
	appName, token := "anyapp", "anytoken"
	s.JwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.ApplicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, &application.GetApplicationNotFound{}).Times(1)

	req, _ := request.New("/api/v1/appadminpolicy/"+appName, request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusForbidden, w.Code)
}

func (s *authzTestSuite) Test_Authorization_AddAdminPolicy_Unauthorized() {
	appName, token := "anyapp", "anytoken"
	s.JwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.ApplicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, &application.GetApplicationUnauthorized{}).Times(1)

	req, _ := request.New("/api/v1/appadminpolicy/"+appName, request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusForbidden, w.Code)
}

func (s *authzTestSuite) Test_Authorization_AddAdminPolicy_Forbidden() {
	appName, token := "anyapp", "anytoken"
	s.JwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.ApplicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, &application.GetApplicationForbidden{}).Times(1)

	req, _ := request.New("/api/v1/appadminpolicy/"+appName, request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusForbidden, w.Code)
}

func (s *authzTestSuite) Test_Authorization_AddAdminPolicy_InternalServerError() {
	appName, token := "anyapp", "anytoken"
	s.JwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.ApplicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, &application.GetApplicationInternalServerError{}).Times(1)

	req, _ := request.New("/api/v1/appadminpolicy/"+appName, request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *authzTestSuite) Test_Authorization_AddAdminPolicy_GenericError() {
	appName, token := "anyapp", "anytoken"
	s.JwtValidator.EXPECT().Validate(token).Return(nil).Times(1)
	s.ApplicationClient.EXPECT().GetApplication(match.GetApplicationRequest(appName), match.GetApplicationAuthRequest(token)).Return(nil, errors.New("any error")).Times(1)

	req, _ := request.New("/api/v1/appadminpolicy/"+appName, request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *authzTestSuite) Test_Authorization_AddAdminPolicy_EndpointWithoutAppName() {
	token := "anytoken"
	s.JwtValidator.EXPECT().Validate(token).Return(nil).Times(1)

	req, _ := request.New("/api/v1/appadminpolicy", request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *authzTestSuite) Test_Authorization_UndefinedPolicy() {
	token := "anytoken"
	s.JwtValidator.EXPECT().Validate(token).Return(nil).Times(1)

	req, _ := request.New("/api/v1/invalidpolicy", request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *authzTestSuite) Test_Authorization_DefaultPolicy_ValidUser() {
	token := "anytoken"
	s.JwtValidator.EXPECT().Validate(token).Return(nil).Times(1)

	req, _ := request.New("/api/v1/defaultpolicy", request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *authzTestSuite) Test_Authorization_DefaultPolicy_AnonymousUser() {
	req, _ := request.New("/api/v1/defaultpolicy")
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}

func (s *authzTestSuite) Test_Authorization_DefaultPolicy_TokenUnauthorized() {
	token := "anytoken"
	s.JwtValidator.EXPECT().Validate(token).Return(apierrors.NewUnauthorizedError()).Times(1)

	req, _ := request.New("/api/v1/defaultpolicy", request.WithBearerAuthorization(token))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusUnauthorized, w.Code)
}
