package authn

type TokenPrincipal interface {
	IsAuthenticated() bool
	Token() string
}
