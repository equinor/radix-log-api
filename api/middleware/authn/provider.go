package authn

import "net/http"

type AuthenticationProvider interface {
	Authenticate(req *http.Request) (user TokenPrincipal, err error)
}
