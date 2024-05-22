package hours01

import (
	"time"
)

func Days(off int, nDay int, choice string) {
	ClearTerminal()
	today := time.Now()
	if nDay != 0 {
		today = time.Date(today.Year(), today.Month(), nDay, 0, 0, 0, 0, today.Location())
	}
	var count int
	for i := 0; i < today.Day(); i++ {
		day := today.AddDate(0, 0, -i)
		if !isWeekend(day.Weekday()) && !isHoliday(day) {
			count++
		}
	}

	finalHours := count*7 - off*7
	FinalPrint(finalHours, choice, today.Day())
}

func isWeekend(day time.Weekday) bool {
	return day == time.Saturday || day == time.Sunday
}

func KnowTheWeekEnd(today int) int {
	var count int
	day := time.Now().AddDate(0, 0, 0)
	if VALUEEND == day.Day() {
		return today
	}
	for i := 1; i < 5; i++ {
		day := time.Now().AddDate(0, 0, i)
		dayName := jourSemaine(day.Weekday())

		if dayName == "samedi" || dayName == "dimanche" {
			return count + today
		} else if VALUEEND == day.Day() {
			return count + today + 1
		} else {
			count++
		}

	}
	return today + count
}
