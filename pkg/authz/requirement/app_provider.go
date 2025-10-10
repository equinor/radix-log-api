package requirement

import (
	"context"

	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/equinor/radix-log-api/pkg/radixapi/client/application"
	"github.com/equinor/radix-log-api/pkg/radixapi/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/rs/zerolog/log"
)

type appProvider struct {
	client application.ClientService
}

func NewApplicationProvider(client application.ClientService) RadixAppProvider {
	return &appProvider{client}
}
func (p *appProvider) GetApplication(ctx context.Context, bearerToken, appName string) (*models.Application, error) {
	ra, err := p.client.GetApplication(
		application.NewGetApplicationParams().WithAppName(appName).WithContext(ctx),
		httptransport.BearerToken(bearerToken))

	if err != nil {
		return nil, err
	}

	if ra.Payload == nil {
		log.Ctx(ctx).Error().Msg("Payload is not set!")
		return nil, apierrors.NewInternalServerError(apierrors.WithCause(ErrMissingRadixAppID))
	}

	return ra.Payload, nil
}
