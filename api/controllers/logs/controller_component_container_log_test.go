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

func Test_ControllerComponentContainerLogSuite(t *testing.T) {
	suite.Run(t, new(controllerComponentContainerLogTestSuite))
}

type controllerComponentContainerLogTestSuite struct {
	controllerTestSuite
}

func (s *controllerComponentContainerLogTestSuite) SetupTest() {
	s.controllerTestSuite.SetupTest()
	s.JwtValidator.EXPECT().Validate(match.IsContext(), gomock.Any()).AnyTimes()
	s.ApplicationClient.EXPECT().GetApplication(gomock.Any(), gomock.Any()).AnyTimes()
}

func (s *controllerComponentContainerLogTestSuite) Test_ContainerLog_Success() {
	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	log := "line1\nline2"
	s.LogService.EXPECT().ComponentContainerLog(match.IsContext(), appName, envName, compName, replicaName, containerId, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.ComponentContainerLogUrl(appName, envName, compName, replicaName, containerId), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Empty(w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *controllerComponentContainerLogTestSuite) Test_ContainerLog_ResponseAsAttachment() {
	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	log := "line1\nline2"
	s.LogService.EXPECT().ComponentContainerLog(match.IsContext(), appName, envName, compName, replicaName, containerId, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.ComponentContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("file", "true")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Equal(`attachment; filename="log.txt"`, w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *controllerComponentContainerLogTestSuite) Test_ContainerLog_WithParams() {
	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	start, end, limit := utils.TimeFormatRFC3339(time.Now()), utils.TimeFormatRFC3339(time.Now().Add(time.Hour)), 500
	s.LogService.EXPECT().ComponentContainerLog(match.IsContext(), appName, envName, compName, replicaName, containerId, &logservice.LogOptions{Timeinterval: &logservice.TimeInterval{Start: start, End: end}, LimitRows: &limit}).Return(bytes.NewReader([]byte{}), nil).Times(1)

	req, _ := request.New(request.ComponentContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("start", start.Format(time.RFC3339)), request.WithQueryParam("end", end.Format(time.RFC3339)), request.WithQueryParam("tail", strconv.Itoa(limit))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *controllerComponentContainerLogTestSuite) Test_ContainerLog_InvalidParam_TailNegative() {
	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	req, _ := request.New(request.ComponentContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("tail", strconv.Itoa(-1))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerComponentContainerLogTestSuite) Test_ContainerLog_InvalidParam_StartNonDate() {
	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	req, _ := request.New(request.ComponentContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("start", "notadate")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerComponentContainerLogTestSuite) Test_ContainerLog_InvalidParam_EndNonDate() {
	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	req, _ := request.New(request.ComponentContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("end", "notadate")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerComponentContainerLogTestSuite) Test_ContainerLog_InvalidParam_FileNonBoolean() {
	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	req, _ := request.New(request.ComponentContainerLogUrl(appName, envName, compName, replicaName, containerId, request.WithQueryParam("file", "notabool")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerComponentContainerLogTestSuite) Test_ContainerLog_LogServiceError() {
	appName, envName, compName, replicaName, containerId := "anyapp", "anyenv", "anycomp", "anyreplica", "anycontainer"
	s.LogService.EXPECT().ComponentContainerLog(match.IsContext(), appName, envName, compName, replicaName, containerId, &logservice.LogOptions{}).Return(bytes.NewReader([]byte{}), errors.New("any error")).Times(1)

	req, _ := request.New(request.ComponentContainerLogUrl(appName, envName, compName, replicaName, containerId), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}
