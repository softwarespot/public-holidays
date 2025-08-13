package holidays

import (
	"strconv"

	datetime "github.com/softwarespot/public-holidays/internal/date-time"
)

func dk(year int) ([]Holiday, error) {
	strYear := strconv.Itoa(year)

	easter := calculateCatholicEaster(year)

	// Specification: https://en.wikipedia.org/wiki/Public_holidays_in_Denmark
	return []Holiday{
		newHoliday(strYear+"-01-01", "Nytårsdag", "New Year's Day"),
		newHoliday(strYear+"-04-18", "Skærtorsdag", "Maundy Thursday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, -2)), "Langfredag", "Good Friday"),
		newHoliday(datetime.ToDateString(easter), "Påskedag", "Easter Sunday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 1)), "Anden påskedag", "Easter Monday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 40)), "Kristi himmelfartsdag", "Ascension Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 49)), "Pinsedag", "Whit Sunday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 50)), "Anden pinsedag", "Whit Monday"),
		newHoliday(strYear+"-06-05", "Grundlovsdag", "Constitution Day"),
		newHoliday(strYear+"-12-25", "Juledag", "Christmas Day"),
		newHoliday(strYear+"-12-26", "Anden juledag", "Boxing Day"),
	}, nil
}
