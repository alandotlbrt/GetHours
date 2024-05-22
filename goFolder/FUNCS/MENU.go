package hours01

import (
	"fmt"
	"os"
	"time"
)

func MAINMENU(language string, path string) {
	DefineTrad(language, path)
	ClearTerminal()
	Menu(path)

}

func Menu(path string) {
	VALUEEND = getLastDayOfMonth()
	fmt.Println(Traductions["option_prompt"])
	fmt.Println(Traductions["Know_Day"])
	fmt.Println(Traductions["Know_Week"])
	fmt.Println(Traductions["Know_NDay"])
	fmt.Println(Traductions["change_Language"])
	fmt.Println(Traductions["exit"])
	var response string
	fmt.Scanln(&response)
	switch response {
	case "1":
		ClearTerminal()
		Days(AskOff(time.Now().Day(), "day"))
	case "2":
		ClearTerminal()
		Days(AskOff(KnowTheWeekEnd(time.Now().Day()), "week"))
	case "3":
		ClearTerminal()
		Days(AskOff(askDays(), "nDay"))
	case "4":
		ClearTerminal()
		ChangeLanguage(path)
	case "5":
		ClearTerminal()
		os.Exit(0)
	}
}

func ChangeLanguage(path string) {
	var language string
	fmt.Println(Traductions["language_ask"])
	fmt.Scanln(&language)
	if language == "2" {
		MAINMENU("fr", path)
	} else if language == "1" {
		MAINMENU("en", path)
	} else {
		MAINMENU("en", path)
	}
}
