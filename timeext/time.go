package timeext

import "time"

// NowUnix returns the current time in Unix seconds
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowUnixMilli returns the current time in Unix milliseconds
func NowUnixMilli() int64 {
	return time.Now().UnixMilli()
}

// AddDate returns the time corresponding to adding the given number of years, months, and days to t.
func AddDate(years int, months int, days int) time.Time {
	return time.Now().AddDate(years, months, days)
}

// SecondsExpired returns true if the given seconds have expired
func SecondsExpired(seconds int64) bool {
	return NowUnix() > seconds
}

// MillisExpired returns true if the given milliseconds have expired
func MillisExpired(millis int64) bool {
	return NowUnixMilli() > millis
}
