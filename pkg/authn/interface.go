package authn

import "net/http"

type ClaimsPrincipal interface {
	IsAuthenticated() bool
	RawToken() string
	IsInRole(role string) bool
}

type Provider interface {
	Authenticate(req *http.Request) (user ClaimsPrincipal, err error)
}
