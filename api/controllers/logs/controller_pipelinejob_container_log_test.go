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

func Test_ControllerPipelineJobContainerLogSuite(t *testing.T) {
	suite.Run(t, new(controllerPipelineJobContainerLogTestSuite))
}

type controllerPipelineJobContainerLogTestSuite struct {
	controllerTestSuite
}

func (s *controllerPipelineJobContainerLogTestSuite) SetupTest() {
	s.controllerTestSuite.SetupTest()
	s.JwtValidator.EXPECT().Validate(match.IsContext(), gomock.Any()).AnyTimes()
	s.ApplicationClient.EXPECT().GetApplication(gomock.Any(), gomock.Any()).AnyTimes()
}

func (s *controllerPipelineJobContainerLogTestSuite) Test_ContainerLog_Success() {
	appName, pipelineJobName, replicaName, containerId := "anyapp", "anypipelinejob", "anyreplica", "anycontainer"
	log := "line1\nline2"
	s.LogService.EXPECT().PipelineJobContainerLog(match.IsContext(), appName, pipelineJobName, replicaName, containerId, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.PipelineJobContainerLogUrl(appName, pipelineJobName, replicaName, containerId), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Empty(w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *controllerPipelineJobContainerLogTestSuite) Test_ContainerLog_ResponseAsAttachment() {
	appName, pipelineJobName, replicaName, containerId := "anyapp", "anypipelinejob", "anyreplica", "anycontainer"
	log := "line1\nline2"
	s.LogService.EXPECT().PipelineJobContainerLog(match.IsContext(), appName, pipelineJobName, replicaName, containerId, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.PipelineJobContainerLogUrl(appName, pipelineJobName, replicaName, containerId, request.WithQueryParam("file", "true")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Equal(`attachment; filename="log.txt"`, w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *controllerPipelineJobContainerLogTestSuite) Test_ContainerLog_WithParams() {
	appName, pipelineJobName, replicaName, containerId := "anyapp", "anypipelinejob", "anyreplica", "anycontainer"
	start, end, limit := utils.TimeFormatRFC3339(time.Now()), utils.TimeFormatRFC3339(time.Now().Add(time.Hour)), 500
	s.LogService.EXPECT().PipelineJobContainerLog(match.IsContext(), appName, pipelineJobName, replicaName, containerId, &logservice.LogOptions{Timeinterval: &logservice.TimeInterval{Start: start, End: end}, LimitRows: &limit}).Return(bytes.NewReader([]byte{}), nil).Times(1)

	req, _ := request.New(request.PipelineJobContainerLogUrl(appName, pipelineJobName, replicaName, containerId, request.WithQueryParam("start", start.Format(time.RFC3339)), request.WithQueryParam("end", end.Format(time.RFC3339)), request.WithQueryParam("tail", strconv.Itoa(limit))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *controllerPipelineJobContainerLogTestSuite) Test_ContainerLog_InvalidParam_TailNegative() {
	appName, pipelineJobName, replicaName, containerId := "anyapp", "anypipelinejob", "anyreplica", "anycontainer"
	req, _ := request.New(request.PipelineJobContainerLogUrl(appName, pipelineJobName, replicaName, containerId, request.WithQueryParam("tail", strconv.Itoa(-1))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerPipelineJobContainerLogTestSuite) Test_ContainerLog_InvalidParam_StartNonDate() {
	appName, pipelineJobName, replicaName, containerId := "anyapp", "anypipelinejob", "anyreplica", "anycontainer"
	req, _ := request.New(request.PipelineJobContainerLogUrl(appName, pipelineJobName, replicaName, containerId, request.WithQueryParam("start", "notadate")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerPipelineJobContainerLogTestSuite) Test_ContainerLog_InvalidParam_EndNonDate() {
	appName, pipelineJobName, replicaName, containerId := "anyapp", "anypipelinejob", "anyreplica", "anycontainer"
	req, _ := request.New(request.PipelineJobContainerLogUrl(appName, pipelineJobName, replicaName, containerId, request.WithQueryParam("end", "notadate")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerPipelineJobContainerLogTestSuite) Test_ContainerLog_InvalidParam_FileNonBoolean() {
	appName, pipelineJobName, replicaName, containerId := "anyapp", "anypipelinejob", "anyreplica", "anycontainer"
	req, _ := request.New(request.PipelineJobContainerLogUrl(appName, pipelineJobName, replicaName, containerId, request.WithQueryParam("file", "notabool")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *controllerPipelineJobContainerLogTestSuite) Test_ContainerLog_LogServiceError() {
	appName, pipelineJobName, replicaName, containerId := "anyapp", "anypipelinejob", "anyreplica", "anycontainer"
	s.LogService.EXPECT().PipelineJobContainerLog(match.IsContext(), appName, pipelineJobName, replicaName, containerId, &logservice.LogOptions{}).Return(bytes.NewReader([]byte{}), errors.New("any error")).Times(1)

	req, _ := request.New(request.PipelineJobContainerLogUrl(appName, pipelineJobName, replicaName, containerId), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}
