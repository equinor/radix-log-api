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

func Test_LogControllerJobReplicaLogTestSuite(t *testing.T) {
	suite.Run(t, new(logControllerJobReplicaLogTestSuite))
}

type logControllerJobReplicaLogTestSuite struct {
	controllerTestSuite
}

func (s *logControllerJobReplicaLogTestSuite) SetupTest() {
	s.controllerTestSuite.SetupTest()
	s.JwtValidator.EXPECT().Validate(match.IsContext(), gomock.Any()).AnyTimes()
	s.ApplicationClient.EXPECT().GetApplication(gomock.Any(), gomock.Any()).AnyTimes()
}

func (s *logControllerJobReplicaLogTestSuite) Test_ReplicaLog_Success() {
	appName, envName, jobCompName, jobName, replicaName := "anyapp", "anyenv", "anyjobcomp", "anyjob", "anyreplica"
	log := "line1\nline2"
	s.LogService.EXPECT().JobPodLog(match.IsContext(), appName, "some-random-id", envName, jobCompName, jobName, replicaName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.JobReplicaLogUrl(appName, envName, jobCompName, jobName, replicaName), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Empty(w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *logControllerJobReplicaLogTestSuite) Test_ReplicaLog_ResponseAsAttachment() {
	appName, envName, jobCompName, jobName, replicaName := "anyapp", "anyenv", "anyjobcomp", "anyjob", "anyreplica"
	log := "line1\nline2"
	s.LogService.EXPECT().JobPodLog(match.IsContext(), appName, "some-random-id", envName, jobCompName, jobName, replicaName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte(log)), nil).Times(1)

	req, _ := request.New(request.JobReplicaLogUrl(appName, envName, jobCompName, jobName, replicaName, request.WithQueryParam("file", "true")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	s.Equal(`attachment; filename="log.txt"`, w.Header().Get("Content-Disposition"))
	actual, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	s.Equal(log, string(actual))
}

func (s *logControllerJobReplicaLogTestSuite) Test_ReplicaLog_WithParams() {
	appName, envName, jobCompName, jobName, replicaName := "anyapp", "anyenv", "anyjobcomp", "anyjob", "anyreplica"
	start, end, limit := utils.TimeFormatRFC3339(time.Now()), utils.TimeFormatRFC3339(time.Now().Add(time.Hour)), 500
	s.LogService.EXPECT().JobPodLog(match.IsContext(), appName, "some-random-id", envName, jobCompName, jobName, replicaName, &logservice.LogOptions{Timeinterval: &logservice.TimeInterval{Start: start, End: end}, LimitRows: &limit}).Return(bytes.NewReader([]byte{}), nil).Times(1)

	req, _ := request.New(request.JobReplicaLogUrl(appName, envName, jobCompName, jobName, replicaName, request.WithQueryParam("start", start.Format(time.RFC3339)), request.WithQueryParam("end", end.Format(time.RFC3339)), request.WithQueryParam("tail", strconv.Itoa(limit))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *logControllerJobReplicaLogTestSuite) Test_ReplicaLog_InvalidParam_TailNegative() {
	appName, envName, jobCompName, jobName, replicaName := "anyapp", "anyenv", "anyjobcomp", "anyjob", "anyreplica"
	req, _ := request.New(request.JobReplicaLogUrl(appName, envName, jobCompName, jobName, replicaName, request.WithQueryParam("tail", strconv.Itoa(-1))), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerJobReplicaLogTestSuite) Test_ReplicaLog_InvalidParam_StartNonDate() {
	appName, envName, jobCompName, jobName, replicaName := "anyapp", "anyenv", "anyjobcomp", "anyjob", "anyreplica"
	req, _ := request.New(request.JobReplicaLogUrl(appName, envName, jobCompName, jobName, replicaName, request.WithQueryParam("start", "notadate")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerJobReplicaLogTestSuite) Test_ReplicaLog_InvalidParam_EndNonDate() {
	appName, envName, jobCompName, jobName, replicaName := "anyapp", "anyenv", "anyjobcomp", "anyjob", "anyreplica"
	req, _ := request.New(request.JobReplicaLogUrl(appName, envName, jobCompName, jobName, replicaName, request.WithQueryParam("end", "notadate")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerJobReplicaLogTestSuite) Test_ReplicaLog_InvalidParam_FileNonBoolean() {
	appName, envName, jobCompName, jobName, replicaName := "anyapp", "anyenv", "anyjobcomp", "anyjob", "anyreplica"
	req, _ := request.New(request.JobReplicaLogUrl(appName, envName, jobCompName, jobName, replicaName, request.WithQueryParam("file", "notabool")), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerJobReplicaLogTestSuite) Test_ReplicaLog_LogServiceError() {
	appName, envName, jobCompName, jobName, replicaName := "anyapp", "anyenv", "anyjobcomp", "anyjob", "anyreplica"
	s.LogService.EXPECT().JobPodLog(match.IsContext(), appName, "some-random-id", envName, jobCompName, jobName, replicaName, &logservice.LogOptions{}).Return(bytes.NewReader([]byte{}), errors.New("any error")).Times(1)

	req, _ := request.New(request.JobReplicaLogUrl(appName, envName, jobCompName, jobName, replicaName), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}
