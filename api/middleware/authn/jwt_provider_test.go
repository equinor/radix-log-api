package authn

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func Test_JwtTestSuite(t *testing.T) {
	suite.Run(t, new(jwtTestSuite))
}

type jwtTestSuite struct {
	suite.Suite
	validator *MockJwtValidator
}

func (s *jwtTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.validator = NewMockJwtValidator(ctrl)
}

func (s *jwtTestSuite) Test_MissingAuthorizationHeader() {
	sut := NewJwtProvider(s.validator)
	user, err := sut.Authenticate(&http.Request{})
	s.NoError(err)
	s.Nil(user)
}

func (s *jwtTestSuite) Test_EmptyAuthorizationHeader() {
	sut := NewJwtProvider(s.validator)
	user, err := sut.Authenticate(&http.Request{Header: http.Header{"Authorization": []string{}}})
	s.NoError(err)
	s.Nil(user)
}

func (s *jwtTestSuite) Test_AuthorizationHeaderWithoutBearer() {
	sut := NewJwtProvider(s.validator)
	user, err := sut.Authenticate(&http.Request{Header: http.Header{"Authorization": []string{"foo"}}})
	s.NoError(err)
	s.Nil(user)
}

func (s *jwtTestSuite) Test_AuthorizationHeaderBearerWithoutToken() {
	sut := NewJwtProvider(s.validator)
	user, err := sut.Authenticate(&http.Request{Header: http.Header{"Authorization": []string{"Bearer "}}})
	s.NoError(err)
	s.Nil(user)
}

func (s *jwtTestSuite) Test_AuthorizationHeaderBearerWithToken() {
	token := "anytoken"
	s.validator.EXPECT().Validate(gomock.Any(), token).Return(nil).Times(1)
	sut := NewJwtProvider(s.validator)
	user, err := sut.Authenticate(&http.Request{Header: http.Header{"Authorization": []string{"Bearer " + token}}})
	s.NoError(err)
	s.NotNil(user)
	s.True(user.IsAuthenticated())
	s.Equal(token, user.Token())
}

func (s *jwtTestSuite) Test_AuthorizationHeaderBearerWithToken_ValidationError() {
	errMsg := "any err"
	s.validator.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(errors.New(errMsg)).Times(1)
	sut := NewJwtProvider(s.validator)
	user, err := sut.Authenticate(&http.Request{Header: http.Header{"Authorization": []string{"Bearer anytoken"}}})
	s.ErrorContains(err, errMsg)
	s.Nil(user)
}
