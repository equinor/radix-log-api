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
	end := o.End
	if end.IsZero() {
		end = time.Now()
	}
	return azquery.NewTimeInterval(o.Start, end)
}

type LogOptions struct {
	Timeinterval *TimeInterval
	LimitRows    *int
}

type InventoryOptions struct {
	Timeinterval *TimeInterval
}
