package holidays

import "time"

type handlerFunc func(year int) ([]Holiday, error)

// Taken from URL: http://stackoverflow.com/questions/2510383/how-can-i-calculate-what-date-good-friday-falls-on-given-a-year
func calculateCatholicEaster(y int) time.Time {
	g := y % 19
	c := y / 100
	h := (c - c/4 - (8*c+13)/25 + 19*g + 15) % 30
	i := h - h/28*(1-h/28*(29/(h+1))*((21-g)/11))

	day := i - (y+y/4+i+2-c+c/4)%7 + 28
	month := 3

	if day > 31 {
		month++
		day -= 31
	}
	return time.Date(y, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func locateDay(year int, month time.Month, day int, targetWeekday time.Weekday) time.Time {
	start := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	if start.Weekday() == targetWeekday {
		return start
	}

	daysUntilTarget := int(targetWeekday - start.Weekday())
	if daysUntilTarget < 0 {
		daysUntilTarget += 7
	}
	return start.AddDate(0, 0, daysUntilTarget)
}
