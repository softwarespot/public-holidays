package holidays

import (
	"strconv"
	"time"

	datetime "github.com/softwarespot/public-holidays/internal/date-time"
)

func fi(year int) ([]Holiday, error) {
	strYear := strconv.Itoa(year)

	easter := calculateCatholicEaster(year)
	midsummer := findNextWeekday(year, time.June, 20, time.Saturday)
	allSaints := findNextWeekday(year, time.October, 31, time.Saturday)

	// Specificiation: https://en.wikipedia.org/wiki/Public_holidays_in_Finland
	return []Holiday{
		newHoliday(strYear+"-01-01", "Uudenvuodenpäivä", "New Year's Day"),
		newHoliday(strYear+"-01-06", "Loppiainen", "Epiphany"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, -2)), "Pitkäperjantai", "Good Friday"),
		newHoliday(datetime.ToDateString(easter), "Pääsiäispäivä", "Easter Sunday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 1)), "2. pääsiäispäivä", "Easter Monday"),
		newHoliday(strYear+"-05-01", "Vappu", "May Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 39)), "Helatorstai", "Ascension Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 49)), "Helluntaipäivä", "Whit Sunday"),
		newHoliday(datetime.ToDateString(midsummer.AddDate(0, 0, -1)), "Juhannusaatto", "Midsummer Eve"),
		newHoliday(datetime.ToDateString(midsummer), "Juhannuspäivä", "Midsummer Day"),
		newHoliday(datetime.ToDateString(allSaints), "Pyhäinpäivä", "All Saints' Day"),
		newHoliday(strYear+"-12-06", "Itsenäisyyspäivä", "Independence Day"),
		newHoliday(strYear+"-12-24", "Jouluaatto", "Christmas Eve"),
		newHoliday(strYear+"-12-25", "Joulupäivä", "Christmas Day"),
		newHoliday(strYear+"-12-26", "2. joulupäivä", "Boxing Day"),
	}, nil
}
