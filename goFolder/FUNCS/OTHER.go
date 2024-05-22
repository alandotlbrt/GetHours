package hours01

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func jourSemaine(day time.Weekday) string {
	switch day {
	case time.Monday:
		return "lundi"
	case time.Tuesday:
		return "mardi"
	case time.Wednesday:
		return "mercredi"
	case time.Thursday:
		return "jeudi"
	case time.Friday:
		return "vendredi"
	case time.Saturday:
		return "samedi"
	case time.Sunday:
		return "dimanche"
	default:
		return "jour inconnu"
	}
}

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func isInteger(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
