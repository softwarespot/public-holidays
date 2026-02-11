package testhelpers

import "time"

const layoutDateTime = "2006-01-02 15:04:05"

// ParseAsDateTime parses a string representation of date and time
// in the format "YYYY-MM-DD HH:MM:SS" into a time.Time value
// based on the local time zone
func ParseAsDateTime(tt string) time.Time {
	t, err := time.ParseInLocation(layoutDateTime, tt, time.Local)
	if err != nil {
		return time.Time{}
	}
	return t
}
