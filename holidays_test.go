package holidays_test

import (
	"testing"
	"time"

	holidays "github.com/thezola/br-holidays"
)

func TestIsBrazilianHoliday(t *testing.T) {
	tests := []struct {
		date        time.Time
		isHoliday   bool
		description string
	}{
		{time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), true, "New Year's Day"},
		{time.Date(2023, time.February, 20, 0, 0, 0, 0, time.UTC), true, "Carnaval Monday"},
		{time.Date(2023, time.February, 21, 0, 0, 0, 0, time.UTC), true, "Carnaval Tuesday"},
		{time.Date(2023, time.April, 7, 0, 0, 0, 0, time.UTC), true, "Passion Friday"},
		{time.Date(2023, 4, 9, 0, 0, 0, 0, time.UTC), true, "Easter Sunday"},
		{time.Date(2023, time.April, 21, 0, 0, 0, 0, time.UTC), true, "Tiradentes Day"},
		{time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC), true, "Worker Day"},
		{time.Date(2023, time.June, 8, 0, 0, 0, 0, time.UTC), true, "Corpus Christi"},
		{time.Date(2023, time.September, 7, 0, 0, 0, 0, time.UTC), true, "Independence Day"},
		{time.Date(2023, time.October, 12, 0, 0, 0, 0, time.UTC), true, "Our Lady of Aparecida"},
		{time.Date(2023, time.November, 2, 0, 0, 0, 0, time.UTC), true, "All Souls' Day"},
		{time.Date(2023, time.November, 15, 0, 0, 0, 0, time.UTC), true, "Proclamation of the Republic"},
		{time.Date(2023, time.December, 25, 0, 0, 0, 0, time.UTC), true, "Christmas Day"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := holidays.IsBrazilianHoliday(test.date)
			if result != test.isHoliday {
				t.Errorf("Expected IsBrazilianHoliday(%v) to be %v, but got %v", test.date, test.isHoliday, result)
			}
		})
	}
}
