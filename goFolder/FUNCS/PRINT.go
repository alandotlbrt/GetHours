package hours01

import (
	"fmt"
	"time"
)

func FinalPrint(hours int, choice string, value int) {
	var message string
	if hours < 0 {
		fmt.Println((Traductions["final_print_negative"]))
		return
	}

	if choice == "week" {
		message = fmt.Sprintf(Traductions["final_print_week"], hours)
	} else if choice == "day" {
		message = fmt.Sprintf(Traductions["final_print_day"], hours)
	} else {
		now := time.Now()
		year, month, _ := now.Date()
		date := time.Date(year, month, value, 0, 0, 0, 0, time.UTC)
		message = fmt.Sprintf(Traductions["final_print_nday"], date.Weekday(), date.Month(), date.Day(), hours)
	}
	bar := "H"
	for i := 0; i < len(message)-1; i++ {
		if message[i] != ' ' {
			bar += "="
		}

	}
	bar += "H"
	fmt.Println(bar + "\n" + message + "\n" + bar)
}
