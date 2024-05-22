package hours01

import (
	"fmt"
	"strconv"
	"time"
)

func AskOff(dateEnd int, choice string) (int, int, string) {
	now := time.Now()
	year, month, _ := now.Date()
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(year, month, dateEnd, 0, 0, 0, 0, time.UTC)
	message := fmt.Sprintf(Traductions["days_off_question"], date.Weekday(), date.Month(), date.Day(), date2.Weekday(), date2.Month(), date2.Day())
	fmt.Println(message)
	var response string
	fmt.Scanln(&response)
	if response == "Y" || response == "y" || response == "O" || response == "o" {
		ClearTerminal()
		return (Scan()), dateEnd, choice
	} else if response == "N" || response == "n" {
		ClearTerminal()
		return 0, dateEnd, choice
	} else {
		fmt.Println(Traductions["invalid_input_message"])
	}
	return 100, 100, "100"
}

func askDays() int {
	for {
		ClearTerminal()
		fmt.Println(Traductions["how_many_hours"])
		var input string
		fmt.Scanln(&input)
		if isInteger(input) {
			value, _ := strconv.Atoi(input)
			if !(value > VALUEEND || value < 0) {
				now := time.Now()
				year, month, _ := now.Date()
				date := time.Date(year, month, value, 0, 0, 0, 0, time.UTC)
				ClearTerminal()
				fmt.Printf(Traductions["hours_day_question"], date.Weekday(), date.Month(), date.Day())
				fmt.Scanln(&input)
				for {
					if input == "Y" || input == "y" || input == "O" || input == "o" {
						ClearTerminal()
						return value
					} else if input == "N" || input == "n" {
						break
					}
				}
			} else {
				fmt.Println()
				fmt.Println(Traductions["invalid_days_number"], VALUEEND)
			}
		} else {
			fmt.Println()
			fmt.Println(Traductions["invalid_days_input_number"])
		}
	}

}

func getLastDayOfMonth() int {
	now := time.Now()
	year, month, _ := now.Date()
	nextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := nextMonth.AddDate(0, 0, -1).Day()
	return lastDayOfMonth
}

func Scan() int {
	var input string
	for {
		fmt.Println(Traductions["days_off_prompt"])
		fmt.Scanln(&input)
		if isInteger(input) {
			value, _ := strconv.Atoi(input)
			if !(value > 31 || value < 0) {
				return value
			} else {
				fmt.Println()
				fmt.Println(Traductions["invalid_days_number"], VALUEEND)
			}
		} else {
			fmt.Println()
			fmt.Println(Traductions["invalid_days_input_number"], VALUEEND)
		}

	}
}
