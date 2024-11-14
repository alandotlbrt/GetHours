package main

import (
	"fmt"
	hours01 "hours01/goFolder/FUNCS"
	"os/exec"
	"path/filepath"
)

func BuildExe() {
	hours01.ClearTerminal()
	var nameFile string
	sentence := "     Write the name of the exe you want to build \033[1m(exemple : KnowMyHoursForTheWeekBecauseImBusyToCalculateIt)\033[0m"
	lignePrint := equal(sentence, 0)
	fmt.Println(lignePrint+"\n"+sentence, "\n"+lignePrint)
	fmt.Print("-->")
	fmt.Scanf("%s", &nameFile)

	fichier := "./goFolder/MAIN/getHours.go"
	cmd := exec.Command("go", "build", "-o", nameFile, fichier)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erreur lors de la construction de l'exécutable de GetHours: %v\n", err)
		return
	}
	if nameFile == "" {
		nameFile = "GetHours"
	}
	hours01.ClearTerminal()
	fmt.Println(equal(nameFile, 30))
	fmt.Printf(" %s has been successfully built\n", nameFile)
	fmt.Println(equal(nameFile, 30))
}

func getAbsolutePath() {
	absolutePath, err := filepath.Abs(".")
	if err != nil {
		fmt.Println("Erreur lors de la résolution du chemin absolu:", err)
		return
	}
	hours01.ClearTerminal()
	fmt.Println(equal(absolutePath, 24))
	fmt.Println("  Path of your file :\033[1m", absolutePath, "\033[0m  ")
	fmt.Println(equal(absolutePath, 24))

}

func main() {
	hours01.ClearTerminal()
	fmt.Println(equal("", 35))
	fmt.Println("Welcome to the GetHours01 project")
	fmt.Println("1. build the exe")
	fmt.Println("2. get the absolute path")
	fmt.Println("Don't forget to read the readme :) ")
	fmt.Println(equal("", 35))
	var response string
	fmt.Print("-->")
	fmt.Scanf("%s", &response)
	if response == "1" {
		BuildExe()
	}
	if response == "2" {
		getAbsolutePath()
	}

}

func equal(a string, b int) string {
	var lignePrint string
	if a == "" {
		for i := 0; i < b; i++ {
			lignePrint += "="
		}
	} else {
		for i := 0; i < len(a)+b; i++ {
			lignePrint += "="
		}
	}
	return lignePrint
}
