package holidays

import (
	"strconv"
	"time"

	datetime "github.com/softwarespot/public-holidays/internal/date-time"
)

func se(year int) ([]Holiday, error) {
	strYear := strconv.Itoa(year)

	easter := calculateCatholicEaster(year)
	midsummer := findNextWeekday(year, time.June, 20, time.Saturday)
	allSaints := findNextWeekday(year, time.October, 31, time.Saturday)

	// Specificiation: https://en.wikipedia.org/wiki/Public_holidays_in_Sweden
	return []Holiday{
		newHoliday(strYear+"-01-01", "Nyårsdagen", "New Year's Day"),
		newHoliday(strYear+"-01-06", "Trettondedag jul", "Epiphany"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, -2)), "Långfredagen", "Good Friday"),
		newHoliday(datetime.ToDateString(easter), "Påskdagen", "Easter Sunday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 1)), "Annandag påsk", "Easter Monday"),
		newHoliday(strYear+"-05-01", "Första Maj", "May Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 39)), "Kristi himmelsfärds dag", "Ascension Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 49)), "Pingstdagen", "Whit Sunday"),
		newHoliday(strYear+"-06-06", "Sveriges nationaldag", "National Day of Sweden"),
		newHoliday(datetime.ToDateString(midsummer.AddDate(0, 0, -1)), "Midsommarafton", "Midsummer Eve"),
		newHoliday(datetime.ToDateString(midsummer), "Midsommardagen", "Midsummer Day"),
		newHoliday(datetime.ToDateString(allSaints), "Alla helgons dag", "All Saints' Day"),
		newHoliday(strYear+"-12-24", "Julafton", "Christmas Eve"),
		newHoliday(strYear+"-12-25", "Juldagen", "Christmas Day"),
		newHoliday(strYear+"-12-26", "Nyårsafton", "Boxing Day"),
		newHoliday(strYear+"-12-31", "Annandag jul", "New Year's Eve"),
	}, nil
}
