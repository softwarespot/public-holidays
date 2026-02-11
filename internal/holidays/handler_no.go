package holidays

import (
	"strconv"

	datetime "github.com/softwarespot/public-holidays/internal/date-time"
)

func no(year int) ([]Holiday, error) {
	strYear := strconv.Itoa(year)

	easter := calculateCatholicEaster(year)

	// Specificiation: https://en.wikipedia.org/wiki/Public_holidays_in_Norway
	return []Holiday{
		newHoliday(strYear+"-01-01", "Første nyttårsdag", "New Year's Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, -3)), "Skjærtorsdag", "Maundy Thursday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, -2)), "Långfredagen", "Good Friday"),
		newHoliday(datetime.ToDateString(easter), "Første påskedag", "Easter Sunday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 1)), "Andre påskedag", "Easter Monday"),
		newHoliday(strYear+"-05-01", "Første mai", "May Day"),
		newHoliday(strYear+"-05-17", "Grunnlovsdagen", "Constitution Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 39)), "Kristi himmelfartsdag", "Ascension Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 49)), "Første pinsedag", "Whit Sunday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 50)), "Andre pinsedag", "Whit Monday"),
		newHoliday(strYear+"-12-25", "Første juledag", "Christmas Day"),
		newHoliday(strYear+"-12-26", "Andre juledag", "Boxing Day"),
	}, nil
}
