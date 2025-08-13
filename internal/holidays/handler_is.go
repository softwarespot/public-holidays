package holidays

import (
	"strconv"
	"time"

	datetime "github.com/softwarespot/public-holidays/internal/date-time"
)

func is(year int) ([]Holiday, error) {
	strYear := strconv.Itoa(year)

	easter := calculateCatholicEaster(year)
	firstDayOfSummer := findNextWeekday(year, time.April, 19, time.Thursday)
	commerceDay := findNextWeekday(year, time.August, 1, time.Monday)

	return []Holiday{
		newHoliday(strYear+"-01-01", "Nýársdagur", "New Year's Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, -3)), "Skírdagur", "Maundy Thursday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, -2)), "Föstudagurinn langi", "Good Friday"),
		newHoliday(datetime.ToDateString(easter), "Páskadagur", "Easter Sunday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 1)), "Annar í páskum", "Easter Monday"),
		newHoliday(datetime.ToDateString(firstDayOfSummer), "Sumardagurinn fyrsti", "First Day of Summer"),
		newHoliday(strYear+"-05-01", "Verkalýðsdagurinn", "May Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 39)), "Uppstigningardagur", "Ascension Day"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 49)), "Hvítasunnudagur", "Whit Sunday"),
		newHoliday(datetime.ToDateString(easter.AddDate(0, 0, 50)), "Annar í hvítasunnu", "Whit Monday"),
		newHoliday(strYear+"-06-17", "Þjóðhátíðardagurinn", "National Day"),
		newHoliday(datetime.ToDateString(commerceDay), "Frídagur verslunarmanna", "Commerce Day"),
		newHoliday(strYear+"-12-24", "Aðfangadagur", "Christmas Eve"),
		newHoliday(strYear+"-12-25", "Jóladagur", "Christmas Day"),
		newHoliday(strYear+"-12-26", "Annar í jólum", "Boxing Day"),
		newHoliday(strYear+"-12-31", "Gamlársdagur", "New Year's Eve"),
	}, nil
}
