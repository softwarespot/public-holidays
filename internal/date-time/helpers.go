package datetime

import "time"

const (
	// ISO 8601 date format i.e. YYYY-MM-DD
	layoutDateOnly = "2006-01-02"
)

func ToDateString(t time.Time) string {
	if t.IsZero() {
		return "0000-00-00"
	}
	return t.Format(layoutDateOnly)
}
