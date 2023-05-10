package request

import "net/http"

type RequestOption func(*http.Request)

func WithBearerAuthorization(token string) RequestOption {
	return WithAuthorization("Bearer " + token)
}

func WithAuthorization(value string) RequestOption {
	return func(r *http.Request) {
		r.Header["Authorization"] = []string{value}
	}
}

func New(url string, options ...RequestOption) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	for _, o := range options {
		o(req)
	}

	return req, nil

}
