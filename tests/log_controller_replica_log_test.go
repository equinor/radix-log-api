package tests

import (
	"bytes"
	"errors"
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

func Test_LogControllerReplicaLogTest(t *testing.T) {
	suite.Run(t, new(logControllerReplicaLogTestSuite))
}

type logControllerReplicaLogTestSuite struct {
	testSuite
}

func (s *logControllerReplicaLogTestSuite) SetupTest() {
	s.testSuite.SetupTest()
	s.jwtValidator.EXPECT().Validate(gomock.Any()).AnyTimes()
	s.applicationClient.EXPECT().GetApplication(gomock.Any(), gomock.Any()).AnyTimes()
}

func (s *logControllerReplicaLogTestSuite) Test_ReplicaLog_Success() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName := "anyapp", "anyenv", "anycomp", "anyreplica"
	log := "line1\nline2"
	s.logService.EXPECT().ComponentPodLog(appName, envName, compName, replicaName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.ReplicaLogUrl(appName, envName, compName, replicaName), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Empty(w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *logControllerReplicaLogTestSuite) Test_ReplicaLog_ResponseAsAttachment() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName := "anyapp", "anyenv", "anycomp", "anyreplica"
	log := "line1\nline2"
	s.logService.EXPECT().ComponentPodLog(appName, envName, compName, replicaName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.ReplicaLogUrl(appName, envName, compName, replicaName, request.WithQueryParam("file", "true")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Equal(`attachment; filename="log.txt"`, w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *logControllerReplicaLogTestSuite) Test_ReplicaLog_WithParams() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName := "anyapp", "anyenv", "anycomp", "anyreplica"
	start, end, limit := timeFormatRFC3339(time.Now()), timeFormatRFC3339(time.Now().Add(time.Hour)), 500
	s.logService.EXPECT().ComponentPodLog(appName, envName, compName, replicaName, &logservice.LogOptions{Timeinterval: &logservice.TimeInterval{Start: start, End: end}, LimitRows: &limit}).Return(bytes.NewReader([]byte{}), nil).Times(1)

	req, _ := request.New(request.ReplicaLogUrl(appName, envName, compName, replicaName, request.WithQueryParam("start", start.Format(time.RFC3339)), request.WithQueryParam("end", end.Format(time.RFC3339)), request.WithQueryParam("tail", strconv.Itoa(limit))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *logControllerReplicaLogTestSuite) Test_ReplicaLog_InvalidParam_TailNegative() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName := "anyapp", "anyenv", "anycomp", "anyreplica"
	req, _ := request.New(request.ReplicaLogUrl(appName, envName, compName, replicaName, request.WithQueryParam("tail", strconv.Itoa(-1))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerReplicaLogTestSuite) Test_ReplicaLog_InvalidParam_StartNonDate() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName := "anyapp", "anyenv", "anycomp", "anyreplica"
	req, _ := request.New(request.ReplicaLogUrl(appName, envName, compName, replicaName, request.WithQueryParam("start", "notadate")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerReplicaLogTestSuite) Test_ReplicaLog_InvalidParam_EndNonDate() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName := "anyapp", "anyenv", "anycomp", "anyreplica"
	req, _ := request.New(request.ReplicaLogUrl(appName, envName, compName, replicaName, request.WithQueryParam("end", "notadate")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerReplicaLogTestSuite) Test_ReplicaLog_InvalidParam_FileNonBoolean() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName := "anyapp", "anyenv", "anycomp", "anyreplica"
	req, _ := request.New(request.ReplicaLogUrl(appName, envName, compName, replicaName, request.WithQueryParam("file", "notabool")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerReplicaLogTestSuite) Test_ReplicaLog_LogServiceError() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName := "anyapp", "anyenv", "anycomp", "anyreplica"
	s.logService.EXPECT().ComponentPodLog(appName, envName, compName, replicaName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte{}), errors.New("any error")).Times(1)

	req, _ := request.New(request.ReplicaLogUrl(appName, envName, compName, replicaName), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}
