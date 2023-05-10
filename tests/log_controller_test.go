package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/equinor/radix-log-api/api/models"
	"github.com/equinor/radix-log-api/api/router"
	logservice "github.com/equinor/radix-log-api/services/logs"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func Test_LogControllerTest(t *testing.T) {
	suite.Run(t, new(logControllerTestSuite))
}

type logControllerTestSuite struct {
	testSuite
}

func (s *logControllerTestSuite) SetupTest() {
	s.testSuite.SetupTest()
	s.jwtValidator.EXPECT().Validate(gomock.Any()).AnyTimes()
	s.applicationClient.EXPECT().GetApplication(gomock.Any(), gomock.Any()).AnyTimes()
}

func (s *logControllerTestSuite) Test_ComponentInventory_Success() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	inventory := []logservice.Pod{
		{
			Name:              "pod1",
			CreationTimestamp: timeFormatRFC3339(time.Now()),
			Containers: []logservice.Container{
				{Id: "c1", CreationTimestamp: timeFormatRFC3339(time.Now().Add(1 * time.Minute))},
				{Id: "c2", CreationTimestamp: timeFormatRFC3339(time.Now().Add(2 * time.Minute))},
			},
		},
		{
			Name:              "pod2",
			CreationTimestamp: timeFormatRFC3339(time.Now().Add(3 * time.Minute)),
			Containers: []logservice.Container{
				{Id: "c3", CreationTimestamp: timeFormatRFC3339(time.Now().Add(4 * time.Minute))},
			},
		},
	}
	s.logService.EXPECT().ComponentInventory(appName, envName, compName, &logservice.ComponentPodInventoryOptions{}).Return(inventory, nil).Times(1)

	req, _ := newRequest(newComponentInventoryUrl(appName, envName, compName), withBearerAuthorization("anytoken"))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	expected := models.ComponentInventoryResponse{Replicas: []models.Replica{
		{
			Name:              inventory[0].Name,
			CreationTimestamp: inventory[0].CreationTimestamp,
			Containers: []models.Container{
				{Id: inventory[0].Containers[0].Id, CreationTimestamp: inventory[0].Containers[0].CreationTimestamp},
				{Id: inventory[0].Containers[1].Id, CreationTimestamp: inventory[0].Containers[1].CreationTimestamp},
			},
		},
		{
			Name:              inventory[1].Name,
			CreationTimestamp: inventory[1].CreationTimestamp,
			Containers: []models.Container{
				{Id: inventory[1].Containers[0].Id, CreationTimestamp: inventory[1].Containers[0].CreationTimestamp},
			},
		},
	}}
	var actual models.ComponentInventoryResponse
	s.Equal(http.StatusOK, w.Code)
	json.NewDecoder(w.Body).Decode(&actual)
	s.Equal(expected, actual)
}

func (s *logControllerTestSuite) Test_ComponentInventory_WithParams() {
	sut, err := router.New(s.logService, s.jwtValidator, s.applicationClient)
	s.Require().NoError(err)

	appName, envName, compName := "anyapp", "anyenv", "anycomp"
	start, end := timeFormatRFC3339(time.Now()), timeFormatRFC3339(time.Now().Add(time.Hour))
	s.logService.EXPECT().ComponentInventory(appName, envName, compName, &logservice.ComponentPodInventoryOptions{Timeinterval: &logservice.TimeInterval{Start: start, End: end}}).Times(1)

	req, _ := newRequest(
		newComponentInventoryUrl(appName, envName, compName, withQueryParam("start", start.Format(time.RFC3339)), withQueryParam("end", end.Format(time.RFC3339))),
		withBearerAuthorization("anytoken"),
	)
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}
