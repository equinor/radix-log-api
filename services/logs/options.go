package logs

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
)

type TimeInterval struct {
	Start time.Time
	End   time.Time
}

func (o TimeInterval) AzQueryTimeinterval() azquery.TimeInterval {
	return azquery.NewTimeInterval(o.Start, o.End)
}

type LogOptions struct {
	Timeinterval *TimeInterval
	LimitRows    *int
}

type ComponentPodInventoryOptions struct {
	Timeinterval *TimeInterval
}
