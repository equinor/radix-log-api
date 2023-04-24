package authn

type ClaimsPrincipal interface {
	IsAuthenticated() bool
	Token() string
}
