package tests

import (
	"time"
)

func timeFormatRFC3339(t time.Time) time.Time {
	newt, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return newt
}
