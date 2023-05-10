package tests

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/equinor/radix-log-api/tests/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	logService        *mock.MockLogService
	jwtValidator      *mock.MockJwtValidator
	applicationClient *mock.MockRadixApiApplicationClient
}

func (s *testSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.logService = mock.NewMockLogService(ctrl)
	s.jwtValidator = mock.NewMockJwtValidator(ctrl)
	s.applicationClient = mock.NewMockRadixApiApplicationClient(ctrl)
}

type httpRequestOption func(*http.Request)

func withBearerAuthorization(token string) httpRequestOption {
	return withAuthorization("Bearer " + token)
}

func withAuthorization(value string) httpRequestOption {
	return func(r *http.Request) {
		r.Header["Authorization"] = []string{value}
	}
}

func newRequest(url string, options ...httpRequestOption) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	for _, o := range options {
		o(req)
	}

	return req, nil

}

type urlOption func(*url.URL)

func withQueryParam(name, value string) urlOption {
	return func(u *url.URL) {
		q := u.Query()
		q.Add(name, value)
		u.RawQuery = q.Encode()
	}
}

func newComponentInventoryUrl(appName, envName, componentName string, options ...urlOption) string {
	return newUrl(fmt.Sprintf("/api/v1/applications/%s/environments/%s/components/%s", appName, envName, componentName), options...)
}

func newUrl(path string, options ...urlOption) string {
	u := &url.URL{Path: path}

	for _, o := range options {
		o(u)
	}

	return u.String()
}

func timeFormatRFC3339(t time.Time) time.Time {
	newt, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return newt
}
