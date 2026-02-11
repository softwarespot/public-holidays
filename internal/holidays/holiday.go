package holidays

type Holiday struct {
	Date        string `json:"date"`
	Name        string `json:"name"`
	EnglishName string `json:"englishName"`
}

func newHoliday(date, name, englishName string) Holiday {
	return Holiday{
		Date:        date,
		Name:        name,
		EnglishName: englishName,
	}
}
