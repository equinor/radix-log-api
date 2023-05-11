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

func Test_LogControllerContainerLogTest(t *testing.T) {
	suite.Run(t, new(logControllerContainerLogTestSuite))
}

type logControllerContainerLogTestSuite struct {
	testSuite
}

func (s *logControllerContainerLogTestSuite) SetupTest() {
	s.testSuite.SetupTest()
	s.jwtValidator.EXPECT().Validate(gomock.Any()).AnyTimes()
	s.applicationClient.EXPECT().GetApplication(gomock.Any(), gomock.Any()).AnyTimes()
}

func (s *logControllerContainerLogTestSuite) Test_ContainerLog_Success() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	log := "line1\nline2"
	s.logService.EXPECT().ComponentContainerLog(appName, envName, compName, replicaName, containerId, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.ContainerLogUrl(appName, envName, compName, replicaName, containerId), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Empty(w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *logControllerContainerLogTestSuite) Test_ContainerLog_ResponseAsAttachment() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	log := "line1\nline2"
	s.logService.EXPECT().ComponentContainerLog(appName, envName, compName, replicaName, containerId, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.ContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("file", "true")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Equal(`attachment; filename="log.txt"`, w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *logControllerContainerLogTestSuite) Test_ContainerLog_WithParams() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	start, end, limit := timeFormatRFC3339(time.Now()), timeFormatRFC3339(time.Now().Add(time.Hour)), 500
	s.logService.EXPECT().ComponentContainerLog(appName, envName, compName, replicaName, containerId, &logservice.LogOptions{Timeinterval: &logservice.TimeInterval{Start: start, End: end}, LimitRows: &limit}).Return(bytes.NewReader([]byte{}), nil).Times(1)

	req, _ := request.New(request.ContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("start", start.Format(time.RFC3339)), request.WithQueryParam("end", end.Format(time.RFC3339)), request.WithQueryParam("tail", strconv.Itoa(limit))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *logControllerContainerLogTestSuite) Test_ContainerLog_InvalidParam_TailNegative() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	req, _ := request.New(request.ContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("tail", strconv.Itoa(-1))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerContainerLogTestSuite) Test_ContainerLog_InvalidParam_StartNonDate() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	req, _ := request.New(request.ContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("start", "notadate")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerContainerLogTestSuite) Test_ContainerLog_InvalidParam_EndNonDate() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	req, _ := request.New(request.ContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("end", "notadate")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerContainerLogTestSuite) Test_ContainerLog_InvalidParam_FileNonBoolean() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	req, _ := request.New(request.ContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("file", "notabool")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerContainerLogTestSuite) Test_ContainerLog_LogServiceError() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	s.logService.EXPECT().ComponentContainerLog(appName, envName, compName, replicaName, containerId, &logservice.LogOptions{}).Return(bytes.NewReader([]byte{}), errors.New("any error")).Times(1)

	req, _ := request.New(request.ContainerLogUrl(appName, envName, compName, replicaName, containerId), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}
