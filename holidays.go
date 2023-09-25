package holidays

import (
	"fmt"
	"time"
)

func IsBrazilianHoliday(date time.Time) bool {

	holidays := [][]int{
		{1, 1},   // New Year's Day
		{4, 21},  // Tiradentes Day
		{9, 7},   // Independence Day
		{10, 12}, // Our Lady of Aparecida
		{11, 2},  // All Souls' Day
		{11, 15}, // Proclamation of the Republic
		{12, 25}, // Christmas Day
	}

	for _, holiday := range holidays {
		if int(date.Month()) == holiday[0] && date.Day() == holiday[1] {
			return true
		}
	}

	// Check for Carnival
	carnival := time.Date(date.Year(), 2, 1, 0, 0, 0, 0, time.UTC)
	if date.Sub(carnival) < 48*time.Hour {
		return true
	}

	// Check for Good Friday
	goodFriday := calculateGoodFriday(date.Year())
	if goodFriday.Day() == date.Day() && goodFriday.Month() == date.Month() {
		return true
	}

	//Check for Corpus Christi
	corpusChristi := calculateCorpusChristi(date.Year())
	if corpusChristi.Day() == date.Day() && corpusChristi.Month() == date.Month() {
		return true
	}

	return false
}

func calculateGoodFriday(year int) time.Time {
	// Calculate Easter Sunday
	a := year % 19
	b := year % 4
	c := year % 7
	d := (19*a + 24) % 30
	e := (2*b + 4*c + 6*d + 5) % 7
	f := d + e

	// Calculate Good Friday
	goodFriday := time.Date(year, 3, 22+f, 0, 0, 0, 0, time.UTC)
	return goodFriday
}

func calculateCorpusChristi(year int) time.Time {
	// Calculate Good Friday
	goodFriday := calculateGoodFriday(year)

	// Calculate Corpus Christi
	corpusChristi := goodFriday.AddDate(0, 0, 60)
	return corpusChristi
}

func main() {
	year := 2022
	goodFriday := calculateGoodFriday(year)
	corpusChristi := calculateCorpusChristi(year)

	fmt.Printf("Good Friday in %d: %s\n", year, goodFriday.Format("2006-01-02"))
	fmt.Printf("Corpus Christi in %d: %s\n", year, corpusChristi.Format("2006-01-02"))
}
