package tests

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/equinor/radix-log-api/api/router"
	logservice "github.com/equinor/radix-log-api/services/logs"
	"github.com/equinor/radix-log-api/tests/internal/request"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func Test_LogControllerComponentLogTest(t *testing.T) {
	suite.Run(t, new(logControllerComponentLogTestSuite))
}

type logControllerComponentLogTestSuite struct {
	testSuite
}

func (s *logControllerComponentLogTestSuite) SetupTest() {
	s.testSuite.SetupTest()
	s.jwtValidator.EXPECT().Validate(gomock.Any()).AnyTimes()
	s.applicationClient.EXPECT().GetApplication(gomock.Any(), gomock.Any()).AnyTimes()
}

func (s *logControllerComponentLogTestSuite) Test_ComponentInventory_Success() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	log := "line1\nline2"
	s.logService.EXPECT().ComponentLog(appName, envName, compName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.ComponentLogUrl(appName, envName, compName), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Empty(w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *logControllerComponentLogTestSuite) Test_ComponentInventory_WithParams() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	start, end, limit := timeFormatRFC3339(time.Now()), timeFormatRFC3339(time.Now().Add(time.Hour)), 500
	s.logService.EXPECT().ComponentLog(appName, envName, compName, &logservice.LogOptions{Timeinterval: &logservice.TimeInterval{Start: start, End: end}, LimitRows: &limit}).Return(bytes.NewReader([]byte{}), nil).Times(1)

	req, _ := request.New(request.ComponentLogUrl(appName, envName, compName, request.WithQueryParam("start", start.Format(time.RFC3339)), request.WithQueryParam("end", end.Format(time.RFC3339)), request.WithQueryParam("tail", strconv.Itoa(limit))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)

}
