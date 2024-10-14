package holidays

import (
	"testing"
	"time"

	testhelpers "github.com/softwarespot/public-holidays/test-helpers"
)

func Test_findNextWeekday(t *testing.T) {
	tests := []struct {
		name          string
		year          int
		month         time.Month
		day           int
		targetWeekday time.Weekday
		want          time.Time
	}{
		{
			name:          "First Thursday on or after April 19, 2024",
			year:          2024,
			month:         time.April,
			day:           19,
			targetWeekday: time.Thursday,
			want:          testhelpers.ParseAsDateTime("2024-04-25 00:00:00"),
		},
		{
			name:          "First Monday in August 2024",
			year:          2024,
			month:         time.August,
			day:           1,
			targetWeekday: time.Monday,
			want:          testhelpers.ParseAsDateTime("2024-08-05 00:00:00"),
		},
		{
			name:          "When start date is the target weekday",
			year:          2024,
			month:         time.January,
			day:           1,
			targetWeekday: time.Monday,
			want:          testhelpers.ParseAsDateTime("2024-01-01 00:00:00"),
		},
		{
			name:          "When target weekday is earlier in the week",
			year:          2024,
			month:         time.July,
			day:           5,
			targetWeekday: time.Monday,
			want:          testhelpers.ParseAsDateTime("2024-07-08 00:00:00"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findNextWeekday(tt.year, tt.month, tt.day, tt.targetWeekday)
			testhelpers.AssertEqual(t, got, tt.want)
		})
	}
}
