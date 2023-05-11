package logs

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/equinor/radix-log-api/internal/tests/request"
	"github.com/equinor/radix-log-api/internal/tests/utils"
	logservice "github.com/equinor/radix-log-api/services/logs"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func Test_ControllerComponentLogTestSuite(t *testing.T) {
	suite.Run(t, new(controllerComponentLogTestSuite))
}

type controllerComponentLogTestSuite struct {
	controllerTestSuite
}

func (s *controllerComponentLogTestSuite) SetupTest() {
	s.controllerTestSuite.SetupTest()
	s.JwtValidator.EXPECT().Validate(gomock.Any()).AnyTimes()
	s.ApplicationClient.EXPECT().GetApplication(gomock.Any(), gomock.Any()).AnyTimes()
}

func (s *controllerComponentLogTestSuite) Test_ComponentLog_Success() {
	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	log := "line1\nline2"
	s.LogService.EXPECT().ComponentLog(appName, envName, compName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req := s.newRequest(request.ComponentLogUrl(appName, envName, compName))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Empty(w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *controllerComponentLogTestSuite) Test_ComponentLog_ResponseAsAttachment() {
	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	log := "line1\nline2"
	s.LogService.EXPECT().ComponentLog(appName, envName, compName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req := s.newRequest(request.ComponentLogUrl(appName, envName, compName, request.WithQueryParam("file", "true")))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Equal(`attachment; filename="log.txt"`, w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *controllerComponentLogTestSuite) Test_ComponentLog_WithParams() {
	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	start, end, limit := utils.TimeFormatRFC3339(time.Now()), utils.TimeFormatRFC3339(time.Now().Add(time.Hour)), 500
	s.LogService.EXPECT().ComponentLog(appName, envName, compName, &logservice.LogOptions{Timeinterval: &logservice.TimeInterval{Start: start, End: end}, LimitRows: &limit}).Return(bytes.NewReader([]byte{}), nil).Times(1)

	req := s.newRequest(request.ComponentLogUrl(appName, envName, compName, request.WithQueryParam("start", start.Format(time.RFC3339)), request.WithQueryParam("end", end.Format(time.RFC3339)), request.WithQueryParam("tail", strconv.Itoa(limit))))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *controllerComponentLogTestSuite) Test_ComponentLog_InvalidParam_TailNegative() {
	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	req := s.newRequest(request.ComponentLogUrl(appName, envName, compName, request.WithQueryParam("tail", strconv.Itoa(-1))))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerComponentLogTestSuite) Test_ComponentLog_InvalidParam_StartNonDate() {
	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	req := s.newRequest(request.ComponentLogUrl(appName, envName, compName, request.WithQueryParam("start", "notadate")))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerComponentLogTestSuite) Test_ComponentLog_InvalidParam_EndNonDate() {
	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	req := s.newRequest(request.ComponentLogUrl(appName, envName, compName, request.WithQueryParam("end", "notadate")))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerComponentLogTestSuite) Test_ComponentLog_InvalidParam_FileNonBoolean() {
	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	req := s.newRequest(request.ComponentLogUrl(appName, envName, compName, request.WithQueryParam("file", "notabool")))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerComponentLogTestSuite) Test_ComponentLog_LogServiceError() {
	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	s.LogService.EXPECT().ComponentLog(appName, envName, compName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte{}), errors.New("any error")).Times(1)

	req := s.newRequest(request.ComponentLogUrl(appName, envName, compName))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}
