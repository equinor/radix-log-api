package error

import (
	"errors"

	apierrors "github.com/equinor/radix-log-api/api/errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	hasErrors := len(ctx.Errors) > 0
	if !hasErrors {
		return
	}

	apiStatus := apierrors.NewInternalServerError().APIStatus
	for _, e := range ctx.Errors {
		logrus.Error(e)
		apierr := apierrors.APIStatus(nil)
		if ok := errors.As(e, &apierr); ok {
			apiStatus = apierr.Status()
		}
	}

	ctx.JSON(int(apiStatus.Code), apiStatus)
}
