package tests

import (
	"fmt"
	"net/http"
)

type httpRequestOption func(*http.Request)

func withBearerAuthorization(token string) httpRequestOption {
	return withAuthorization("Bearer " + token)
}

func withAuthorization(value string) httpRequestOption {
	return func(r *http.Request) {
		r.Header["Authorization"] = []string{value}
	}
}

func newInventoryRequest(appName, envName, componentName string, options ...httpRequestOption) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/applications/%s/environments/%s/components/%s", appName, envName, componentName), nil)
	if err != nil {
		return nil, err
	}

	for _, o := range options {
		o(req)
	}

	return req, nil
}
