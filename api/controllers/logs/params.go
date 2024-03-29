package logs

import (
	"time"

	logservice "github.com/equinor/radix-log-api/pkg/services/logs"
	"github.com/gin-gonic/gin"
)

type timeIntervalParams struct {
	Start *time.Time `form:"start" time_format:"2006-01-02T15:04:05Z07:00"`
	End   *time.Time `form:"end" time_format:"2006-01-02T15:04:05Z07:00"`
}

type logParams struct {
	timeIntervalParams
	Tail *int `form:"tail" binding:"omitempty,min=0"`
	File bool `form:"file"`
}

// AsLogOptions converts the parameters to LogOptions parameters
func (p *logParams) AsLogOptions() logservice.LogOptions {
	options := logservice.LogOptions{LimitRows: p.Tail}
	if p.Start != nil || p.End != nil {
		timeInverval := logservice.TimeInterval{}
		if p.Start != nil {
			timeInverval.Start = *p.Start
		}
		if p.End != nil {
			timeInverval.End = *p.End
		}
		options.Timeinterval = &timeInverval
	}

	return options
}

type inventoryParams struct {
	timeIntervalParams
}

// AsInventoryOptions converts the parameters to AsInventoryOptions parameters
func (p *inventoryParams) AsInventoryOptions() logservice.InventoryOptions {
	var options logservice.InventoryOptions
	if p.Start != nil || p.End != nil {
		timeInverval := logservice.TimeInterval{}
		if p.Start != nil {
			timeInverval.Start = *p.Start
		}
		if p.End != nil {
			timeInverval.End = *p.End
		}
		options.Timeinterval = &timeInverval
	}

	return options
}

func paramsFromContext[T any](ctx *gin.Context) (*T, error) {
	params := new(T)
	err := ctx.ShouldBindQuery(params)
	return params, err
}
