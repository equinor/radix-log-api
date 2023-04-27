package logs

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
)

type TimeintervalOptions struct {
	Start time.Time
	End   time.Time
}

func (o TimeintervalOptions) AzQueryTimeinterval() azquery.TimeInterval {
	return azquery.NewTimeInterval(o.Start, o.End)
}

type ComponentLogOptions struct {
	Timeinterval *TimeintervalOptions
	LimitRows    *int
}

type ComponentPodInventoryOptions struct {
	Timeinterval *TimeintervalOptions
}
