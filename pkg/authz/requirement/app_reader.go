package requirement

import (
	"context"
	"errors"

	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/api/middleware/authz"
	"github.com/equinor/radix-log-api/api/params"
	"github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	"github.com/equinor/radix-log-api/pkg/radixapi/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const radixAppIDKey = "radixAppId"

var ErrMissingRadixAppID = errors.New("appId is missing from RadixRegistration")

type RadixAppProvider interface {
	GetApplication(ctx context.Context, bearerToken, appName string) (*models.Application, error)
}

type appReaderRequirement struct {
	appProvider RadixAppProvider
}

// NewAppReaderRequirement returns a requirement that is satisfied when the caller can read the application
// through Radix API. Radix API authorizes GET /applications/{appName} using the caller's bearer token, and
// application readers are intentionally allowed to view logs (the radix-app-reader-envs ClusterRole in
// radix-operator grants pods/log). A successful lookup is therefore sufficient; no admin check is intended.
func NewAppReaderRequirement(appProvider RadixAppProvider) authz.Requirement {
	return &appReaderRequirement{
		appProvider: appProvider,
	}
}

func (r *appReaderRequirement) ValidateRequirement(ctx *authz.AuthorizationContext) error {
	var params params.App
	if err := ctx.GinCtx().ShouldBindUri(&params); err != nil {
		return apierrors.NewInternalServerError(apierrors.WithCause(err))
	}

	ra, err := r.appProvider.GetApplication(ctx.GinCtx().Request.Context(), ctx.User().Token(), params.AppName)

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

	if ra == nil || ra.Registration == nil || ra.Registration.AppID == nil || *ra.Registration.AppID == "" {
		return apierrors.NewInternalServerError(apierrors.WithCause(ErrMissingRadixAppID))
	}

	ctx.GinCtx().Set(radixAppIDKey, *ra.Registration.AppID)
	return nil
}

func GetAppId(ctx *gin.Context) (string, error) {
	appId := ctx.GetString(radixAppIDKey)
	if len(appId) == 0 {
		log.Error().Msg("You cannot call GetAppId without using the paired AppReader requirement!")
		return "", apierrors.NewForbiddenError(apierrors.WithCause(ErrMissingRadixAppID))
	}

	return appId, nil
}
