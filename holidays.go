package holidays

import (
	"time"
)

func IsBrazilianHoliday(date time.Time) bool {

	holidays := [][]int{
		{1, 1},   // New Year's Day
		{4, 21},  // Tiradentes Day
		{5, 1},   // Worker Day
		{9, 7},   // Independence Day
		{10, 12}, // Our Lady of Aparecida
		{11, 2},  // All Souls' Day
		{11, 15}, // Proclamation of the Republic
		{11, 20}, // Black Consciousness
		{12, 25}, // Christmas Day
	}

	for _, holiday := range holidays {
		if int(date.Month()) == holiday[0] && date.Day() == holiday[1] {
			return true
		}
	}

	//Check for Easter Sunday - this is the base of the other christian holidays
	easterSunday := calculateEasterSunday(date.Year())
	if date.Day() == easterSunday.Day() && date.Month() == easterSunday.Month() {
		return true
	}

	// Check for Carnaval
	//there are more factors involved, some brazilians take the whole week, im just checking for monday and tuesday here.
	carnavalTuesday := easterSunday.AddDate(0, 0, -47)
	if date.Day() == carnavalTuesday.Day() && date.Month() == carnavalTuesday.Month() {
		return true
	}
	carnavalMonday := carnavalTuesday.AddDate(0, 0, -1)
	if date.Day() == carnavalMonday.Day() && date.Month() == carnavalMonday.Month() {
		return true
	}

	// Check for Good Friday - 2 days before easter sunday
	goodFriday := easterSunday.AddDate(0, 0, -2)
	if goodFriday.Day() == date.Day() && goodFriday.Month() == date.Month() {
		return true
	}

	//Check for Corpus Christi - 60 days after easter sunday
	corpusChristi := easterSunday.AddDate(0, 0, 60)
	if corpusChristi.Day() == date.Day() && corpusChristi.Month() == date.Month() {
		return true
	}

	return false
}

func calculateEasterSunday(year int) time.Time {
	// Calculating easter sunday using the computus algorithm instead of gauss algorithm
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451
	month := (h + l - 7*m + 114) / 31
	day := (h + l - 7*m + 114) % 31

	return time.Date(year, time.Month(month), day+1, 0, 0, 0, 0, time.UTC)

}
