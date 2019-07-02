package utils

import (
	"time"
)

// DateConvertUnix ...
func DateConvertUnix(d int64, format string) string {
	unixTimeUTC := time.Unix((d / 1000), 0) //gives unix time stamp in utc
	switch format {
	case time.RFC3339:
		unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format
		return unitTimeInRFC3339
	default:
		return unixTimeUTC.Format("2006-01-02")
	}
}
