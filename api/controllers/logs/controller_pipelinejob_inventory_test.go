package logs

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/equinor/radix-log-api/api/models"
	"github.com/equinor/radix-log-api/internal/tests/match"
	"github.com/equinor/radix-log-api/internal/tests/request"
	"github.com/equinor/radix-log-api/internal/tests/utils"
	logservice "github.com/equinor/radix-log-api/pkg/services/logs"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func Test_LogControllerPipelineJobInventoryTestSuite(t *testing.T) {
	suite.Run(t, new(logControllerPipelineJobInventoryTestSuite))
}

type logControllerPipelineJobInventoryTestSuite struct {
	controllerTestSuite
}

func (s *logControllerPipelineJobInventoryTestSuite) SetupTest() {
	s.controllerTestSuite.SetupTest()
	s.JwtValidator.EXPECT().Validate(match.IsContext(), gomock.Any()).AnyTimes()
	s.ApplicationClient.EXPECT().GetApplication(gomock.Any(), gomock.Any()).AnyTimes()
}

func (s *logControllerPipelineJobInventoryTestSuite) Test_PipelineJobInventory_Success() {
	appName, pipelineJobName := "anyapp", "anypipelinejob"
	pod1BaseTime, pod2BaseTime := time.Now(), time.Now().Add(1*time.Hour)
	inventory := []logservice.Pod{
		{
			Name:              "pod1",
			CreationTimestamp: utils.TimeFormatRFC3339(pod1BaseTime),
			LastKnown:         utils.TimeFormatRFC3339(pod1BaseTime.Add(1 * time.Minute)),
			Containers: []logservice.Container{
				{Id: "c1", LastKnown: utils.TimeFormatRFC3339(pod1BaseTime.Add(2 * time.Minute)), CreationTimestamp: utils.TimeFormatRFC3339(utils.TimeFormatRFC3339(pod1BaseTime.Add(3 * time.Minute)).Add(3 * time.Minute))},
				{Id: "c2", LastKnown: utils.TimeFormatRFC3339(time.Now().Add(4 * time.Minute)), CreationTimestamp: utils.TimeFormatRFC3339(time.Now().Add(5 * time.Minute))},
			},
		},
		{
			Name:              "pod2",
			CreationTimestamp: utils.TimeFormatRFC3339(pod2BaseTime),
			LastKnown:         utils.TimeFormatRFC3339(pod2BaseTime.Add(1 * time.Minute)),
			Containers: []logservice.Container{
				{Id: "c3", LastKnown: utils.TimeFormatRFC3339(pod2BaseTime.Add(2 * time.Minute)), CreationTimestamp: utils.TimeFormatRFC3339(utils.TimeFormatRFC3339(pod2BaseTime.Add(3 * time.Minute)).Add(3 * time.Minute))},
			},
		},
	}
	s.LogService.EXPECT().PipelineJobInventory(match.IsContext(), appName, pipelineJobName, &logservice.InventoryOptions{}).Return(inventory, nil).Times(1)

	req, _ := request.New(request.PipelineJobInventoryUrl(appName, pipelineJobName), request.WithBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	expected := models.InventoryResponse{Replicas: []models.Replica{
		{
			Name:              inventory[0].Name,
			CreationTimestamp: inventory[0].CreationTimestamp,
			LastKnown:         inventory[0].LastKnown,
			Containers: []models.Container{
				{Id: inventory[0].Containers[0].Id, LastKnown: inventory[0].Containers[0].LastKnown, CreationTimestamp: inventory[0].Containers[0].CreationTimestamp},
				{Id: inventory[0].Containers[1].Id, LastKnown: inventory[0].Containers[1].LastKnown, CreationTimestamp: inventory[0].Containers[1].CreationTimestamp},
			},
		},
		{
			Name:              inventory[1].Name,
			CreationTimestamp: inventory[1].CreationTimestamp,
			LastKnown:         inventory[1].LastKnown,
			Containers: []models.Container{
				{Id: inventory[1].Containers[0].Id, LastKnown: inventory[1].Containers[0].LastKnown, CreationTimestamp: inventory[1].Containers[0].CreationTimestamp},
			},
		},
	}}
	var actual models.InventoryResponse
	s.Equal(http.StatusOK, w.Code)
	_ = json.NewDecoder(w.Body).Decode(&actual)
	s.Equal(expected, actual)
}

func (s *logControllerPipelineJobInventoryTestSuite) Test_PipelineJobInventory_WithParams() {
	appName, pipelineJobName := "anyapp", "anypipelinejob"
	start, end := utils.TimeFormatRFC3339(time.Now()), utils.TimeFormatRFC3339(time.Now().Add(time.Hour))
	s.LogService.EXPECT().PipelineJobInventory(match.IsContext(), appName, pipelineJobName, &logservice.InventoryOptions{Timeinterval: &logservice.TimeInterval{Start: start, End: end}}).Times(1)

	req, _ := request.New(
		request.PipelineJobInventoryUrl(appName, pipelineJobName, request.WithQueryParam("start", start.Format(time.RFC3339)), request.WithQueryParam("end", end.Format(time.RFC3339))),
		request.WithBearerAuthorization("anytoken"),
	)
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *logControllerPipelineJobInventoryTestSuite) Test_PipelineJobInventory_InvalidParam_StartNonData() {
	appName, pipelineJobName := "anyapp", "anypipelinejob"
	req, _ := request.New(
		request.PipelineJobInventoryUrl(appName, pipelineJobName, request.WithQueryParam("start", "notadate")),
		request.WithBearerAuthorization("anytoken"),
	)
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerPipelineJobInventoryTestSuite) Test_PipelineJobInventory_InvalidParam_EndNonDate() {
	appName, pipelineJobName := "anyapp", "anypipelinejob"
	req, _ := request.New(
		request.PipelineJobInventoryUrl(appName, pipelineJobName, request.WithQueryParam("end", "notadate")),
		request.WithBearerAuthorization("anytoken"),
	)
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *logControllerPipelineJobInventoryTestSuite) Test_PipelineJobInventory_LogServiceError() {
	appName, pipelineJobName := "anyapp", "anypipelinejob"
	s.LogService.EXPECT().PipelineJobInventory(match.IsContext(), appName, pipelineJobName, &logservice.InventoryOptions{}).Return(nil, errors.New("any error")).Times(1)

	req, _ := request.New(
		request.PipelineJobInventoryUrl(appName, pipelineJobName),
		request.WithBearerAuthorization("anytoken"),
	)
	w := httptest.NewRecorder()
	s.sut().ServeHTTP(w, req)
	s.Equal(http.StatusInternalServerError, w.Code)
}