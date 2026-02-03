package match

import (
	"context"

	"go.uber.org/mock/gomock"
)

func IsContext() gomock.Matcher {
	return &contextMatcher{}
}

type contextMatcher struct {
}

func (*contextMatcher) Matches(x interface{}) bool {
	_, ok := x.(context.Context)
	return ok
}

func (*contextMatcher) String() string {
	return "is context"
}
