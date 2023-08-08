package requirement

import (
	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/middleware/authz"
	"github.com/equinor/radix-log-api/api/params"
	"github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	httptransport "github.com/go-openapi/runtime/client"
)

type appOwnerRequirement struct {
	applicationClient application.ClientService
}

func NewAppOwnerRequirement(applicationClient application.ClientService) authz.Requirement {
	return &appOwnerRequirement{
		applicationClient: applicationClient,
	}
}

func (r *appOwnerRequirement) ValidateRequirement(ctx *authz.AuthorizationContext) error {
	var params params.App
	if err := ctx.GinCtx().ShouldBindUri(&params); err != nil {
		return apierrors.NewInternalServerError(apierrors.WithCause(err))
	}

	_, err := r.applicationClient.GetApplication(
		application.NewGetApplicationParams().WithAppName(params.AppName).WithContext(ctx.GinCtx().Request.Context()),
		httptransport.BearerToken(ctx.User().Token()))

	switch err.(type) {
	case *application.GetApplicationUnauthorized:
		return apierrors.NewForbiddenError(apierrors.WithCause(err))
	case *application.GetApplicationForbidden:
		return apierrors.NewForbiddenError(apierrors.WithCause(err))
	case *application.GetApplicationNotFound:
		return apierrors.NewForbiddenError(apierrors.WithCause(err))
	}

	return err
}
