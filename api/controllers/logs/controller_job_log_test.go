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

	"github.com/equinor/radix-log-api/internal/tests/match"
	"github.com/equinor/radix-log-api/internal/tests/request"
	"github.com/equinor/radix-log-api/internal/tests/utils"
	logservice "github.com/equinor/radix-log-api/pkg/services/logs"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func Test_ControllerJobLogTestSuite(t *testing.T) {
	suite.Run(t, new(controllerJobLogTestSuite))
}

type controllerJobLogTestSuite struct {
	controllerTestSuite
}

func (s *controllerJobLogTestSuite) SetupTest() {
	s.controllerTestSuite.SetupTest()
	s.JwtValidator.EXPECT().Validate(match.IsContext(), gomock.Any()).AnyTimes()
	s.ApplicationClient.EXPECT().GetApplication(gomock.Any(), gomock.Any()).AnyTimes()
}

func (s *controllerJobLogTestSuite) Test_JobLog_Success() {
	appName, envName, jobCompName, jobName := "anyapp", "anyenv", "anyjobcomp", "anyjob"
	log := "line1\nline2"
	s.LogService.EXPECT().JobLog(match.IsContext(), appName, envName, jobCompName, jobName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req := s.newRequest(request.JobLogUrl(appName, envName, jobCompName, jobName))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Empty(w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *controllerJobLogTestSuite) Test_JobLog_ResponseAsAttachment() {
	appName, envName, jobCompName, jobName := "anyapp", "anyenv", "anyjobcomp", "anyjob"
	log := "line1\nline2"
	s.LogService.EXPECT().JobLog(match.IsContext(), appName, envName, jobCompName, jobName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req := s.newRequest(request.JobLogUrl(appName, envName, jobCompName, jobName, request.WithQueryParam("file", "true")))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Equal(`attachment; filename="log.txt"`, w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *controllerJobLogTestSuite) Test_JobLog_WithParams() {
	appName, envName, jobCompName, jobName := "anyapp", "anyenv", "anyjobcomp", "anyjob"
	start, end, limit := utils.TimeFormatRFC3339(time.Now()), utils.TimeFormatRFC3339(time.Now().Add(time.Hour)), 500
	s.LogService.EXPECT().JobLog(match.IsContext(), appName, envName, jobCompName, jobName, &logservice.LogOptions{Timeinterval: &logservice.TimeInterval{Start: start, End: end}, LimitRows: &limit}).Return(bytes.NewReader([]byte{}), nil).Times(1)

	req := s.newRequest(request.JobLogUrl(appName, envName, jobCompName, jobName, request.WithQueryParam("start", start.Format(time.RFC3339)), request.WithQueryParam("end", end.Format(time.RFC3339)), request.WithQueryParam("tail", strconv.Itoa(limit))))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *controllerJobLogTestSuite) Test_JobLog_InvalidParam_TailNegative() {
	appName, envName, jobCompName, jobName := "anyapp", "anyenv", "anyjobcomp", "anyjob"
	req := s.newRequest(request.JobLogUrl(appName, envName, jobCompName, jobName, request.WithQueryParam("tail", strconv.Itoa(-1))))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerJobLogTestSuite) Test_JobLog_InvalidParam_StartNonDate() {
	appName, envName, jobCompName, jobName := "anyapp", "anyenv", "anyjobcomp", "anyjob"
	req := s.newRequest(request.JobLogUrl(appName, envName, jobCompName, jobName, request.WithQueryParam("start", "notadate")))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerJobLogTestSuite) Test_JobLog_InvalidParam_EndNonDate() {
	appName, envName, jobCompName, jobName := "anyapp", "anyenv", "anyjobcomp", "anyjob"
	req := s.newRequest(request.JobLogUrl(appName, envName, jobCompName, jobName, request.WithQueryParam("end", "notadate")))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerJobLogTestSuite) Test_JobLog_InvalidParam_FileNonBoolean() {
	appName, envName, jobCompName, jobName := "anyapp", "anyenv", "anyjobcomp", "anyjob"
	req := s.newRequest(request.JobLogUrl(appName, envName, jobCompName, jobName, request.WithQueryParam("file", "notabool")))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerJobLogTestSuite) Test_JobLog_LogServiceError() {
	appName, envName, jobCompName, jobName := "anyapp", "anyenv", "anyjobcomp", "anyjob"
	s.LogService.EXPECT().JobLog(match.IsContext(), appName, envName, jobCompName, jobName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte{}), errors.New("any error")).Times(1)

	req := s.newRequest(request.JobLogUrl(appName, envName, jobCompName, jobName))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}
