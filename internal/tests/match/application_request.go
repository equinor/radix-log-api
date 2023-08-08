package match

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	"github.com/go-openapi/runtime"
	"github.com/golang/mock/gomock"
)

func GetApplicationRequest(appName string) gomock.Matcher {
	return &getApplicationMatcher{appName: appName}
}

type getApplicationMatcher struct {
	appName string
}

// Matches returns whether x is a match.
func (m *getApplicationMatcher) Matches(x interface{}) bool {
	if params, ok := x.(*application.GetApplicationParams); ok {
		return params.AppName == m.appName && IsContext().Matches(params.Context)
	}
	return false
}

// String describes what the matcher matches.
func (m *getApplicationMatcher) String() string {
	return fmt.Sprintf("appName '%s' in request", m.appName)
}

func GetApplicationAuthRequest(token string) gomock.Matcher {
	return &getApplicationAuthMatcher{token: token}
}

type getApplicationAuthMatcher struct {
	token string
}

// Matches returns whether x is a match.
func (m *getApplicationAuthMatcher) Matches(x interface{}) bool {
	if authWriter, ok := x.(runtime.ClientAuthInfoWriter); ok {
		req := fakeClientRequest{}
		authWriter.AuthenticateRequest(&req, nil)
		return req.authorizationHeader == "Bearer "+m.token
	}
	return false
}

// String describes what the matcher matches.
func (m *getApplicationAuthMatcher) String() string {
	return fmt.Sprintf("Bearer '%s' in Authorization header", m.token)
}

type fakeClientRequest struct {
	authorizationHeader string
}

func (r *fakeClientRequest) SetHeaderParam(header string, values ...string) error {
	if header == "Authorization" {
		r.authorizationHeader = ""
		if len(values) == 1 {
			r.authorizationHeader = values[0]
		}
	}
	return nil
}

func (*fakeClientRequest) GetHeaderParams() http.Header { return nil }

func (*fakeClientRequest) SetQueryParam(string, ...string) error { return nil }

func (*fakeClientRequest) SetFormParam(string, ...string) error { return nil }

func (*fakeClientRequest) SetPathParam(string, string) error { return nil }

func (*fakeClientRequest) GetQueryParams() url.Values { return nil }

func (*fakeClientRequest) SetFileParam(string, ...runtime.NamedReadCloser) error { return nil }

func (*fakeClientRequest) SetBodyParam(interface{}) error { return nil }

func (*fakeClientRequest) SetTimeout(time.Duration) error { return nil }

func (*fakeClientRequest) GetMethod() string { return "" }

func (*fakeClientRequest) GetPath() string { return "" }

func (*fakeClientRequest) GetBody() []byte { return nil }

func (*fakeClientRequest) GetBodyParam() interface{} { return nil }

func (*fakeClientRequest) GetFileParam() map[string][]runtime.NamedReadCloser { return nil }
