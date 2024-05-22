package hours01

import "time"

func isHoliday(date time.Time) bool {
	holidays := []time.Time{
		time.Date(date.Year(), time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(date.Year(), time.May, 1, 0, 0, 0, 0, time.UTC),
		time.Date(date.Year(), time.May, 8, 0, 0, 0, 0, time.UTC),
		time.Date(date.Year(), time.July, 14, 0, 0, 0, 0, time.UTC),
		time.Date(date.Year(), time.August, 15, 0, 0, 0, 0, time.UTC),
		time.Date(date.Year(), time.November, 1, 0, 0, 0, 0, time.UTC),
		time.Date(date.Year(), time.November, 11, 0, 0, 0, 0, time.UTC),
		time.Date(date.Year(), time.December, 25, 0, 0, 0, 0, time.UTC),
	}
	algo := calculateEaster(date.Year())
	holidays = append(holidays, algo.AddDate(0, 0, 1), algo.AddDate(0, 0, 39), algo.AddDate(0, 0, 50))
	for _, holiday := range holidays {
		if date.Year() == holiday.Year() && date.Month() == holiday.Month() && date.Day() == holiday.Day() {
			return true
		}
	}
	return false
}

func calculateEaster(year int) time.Time {
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
	day := ((h + l - 7*m + 114) % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
