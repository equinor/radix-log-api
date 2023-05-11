package utils

import "time"

func TimeFormatRFC3339(t time.Time) time.Time {
	newt, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return newt
}
