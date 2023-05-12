package authn

import "net/http"

// AuthenticationProvider provides a method to authenticate a request
type AuthenticationProvider interface {
	// Authenticate tries to build a TokenPrincipal from the HTTP request.
	// Returns a TokenPrincipal with user.IsAuthenticated()=true if the request contains valid user information,
	// otherwise return nil or  TokenPrincipal with user.IsAuthenticated()=false.
	// If an error is returned, the authentication handler abort the request immediately write the error to the gin.Context.
	Authenticate(req *http.Request) (user TokenPrincipal, err error)
}
