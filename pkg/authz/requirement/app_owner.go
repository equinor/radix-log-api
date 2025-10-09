package requirement

import (
	"errors"

	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/middleware/authz"
	"github.com/equinor/radix-log-api/api/params"
	"github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	"github.com/gin-gonic/gin"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/rs/zerolog/log"
)

const radixAppIDKey = "radixAppId"

var ErrMissingRadixAppID = errors.New("appId is missing from RadixRegistration")

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

	ra, err := r.applicationClient.GetApplication(
		application.NewGetApplicationParams().WithAppName(params.AppName).WithContext(ctx.GinCtx().Request.Context()),
		httptransport.BearerToken(ctx.User().Token()))

	if err != nil {
		switch err.(type) {
		case *application.GetApplicationUnauthorized:
			return apierrors.NewForbiddenError(apierrors.WithCause(err))
		case *application.GetApplicationForbidden:
			return apierrors.NewForbiddenError(apierrors.WithCause(err))
		case *application.GetApplicationNotFound:
			return apierrors.NewForbiddenError(apierrors.WithCause(err))
		default:
			return err
		}
	}

	if ra == nil || ra.Payload == nil || ra.Payload.Registration == nil || ra.Payload.Registration.AppID == nil || *ra.Payload.Registration.AppID == "" {
		return apierrors.NewInternalServerError(apierrors.WithCause(ErrMissingRadixAppID))
	}

	ctx.GinCtx().Set(radixAppIDKey, *ra.Payload.Registration.AppID)
	return nil
}

func GetAppId(ctx *gin.Context) (string, error) {
	appId := ctx.GetString(radixAppIDKey)
	if len(appId) == 0 {
		log.Error().Msg("You cannot call GetAppId without using the paired AppOwner requirement!")
		return "", apierrors.NewForbiddenError(apierrors.WithCause(ErrMissingRadixAppID))
	}

	return appId, nil
}
