package hours01

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func chargerFichierDeTraduction(langue string, path string) (tradu, error) {
	var settingsFilePath string
	if path == "" {
		settingsFilePath, _ = filepath.Abs("./language")
	} else {
		settingsFilePath = path + "/language"

	}
	fichierName := fmt.Sprintf("%s/%s.json", settingsFilePath, langue)
	fichier, err := os.Open(fichierName)
	if err != nil {
		return nil, err
	}
	defer fichier.Close()
	traductionsJson := make(tradu)
	scanner := bufio.NewScanner(fichier)
	var contenu []byte
	for scanner.Scan() {
		contenu = append(contenu, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(contenu, &traductionsJson); err != nil {
		return nil, err
	}
	return traductionsJson, nil
}

func getTrad(language string, path string) map[string]string {
	langueUtilisateur := language
	traductions, err := chargerFichierDeTraduction(langueUtilisateur, path)
	if err != nil {
		fmt.Println("error while loading traductions: ", err)
		return nil
	}
	return traductions
}

func DefineTrad(language string, path string) {
	Traductions = getTrad(Getsettings(language, path), path)
}
func Getsettings(langue string, path string) string {
	var settingsFilePath string
	if path == "" {
		settingsFilePath, _ = filepath.Abs("./profile/save_settings.txt")

	} else {
		settingsFilePath = path + "/profile/save_settings.txt"

	}
	if _, err := os.Stat(settingsFilePath); os.IsNotExist(err) {
		file, err := os.Create(settingsFilePath)
		if err != nil {
			return ""
		}
		defer file.Close()
		file.WriteString("language : en\n")
	}
	settingsFile, err := os.OpenFile(settingsFilePath, os.O_RDWR, 0644)
	if err != nil {
		return ""
	}
	defer settingsFile.Close()
	scanner := bufio.NewScanner(settingsFile)
	for scanner.Scan() {
		line := scanner.Text()
		if langue == "" {
			if line != "" && strings.HasPrefix(line, "language") {
				language := strings.TrimSpace(line[10:])
				return language
			}
		} else {
			if line != "" && strings.HasPrefix(line, "language") {
				_, err := settingsFile.Seek(0, 0)
				if err != nil {
					return ""
				}
				newLanguageLine := fmt.Sprintf("language : %s\n", langue)
				_, err = settingsFile.WriteString(newLanguageLine)
				if err != nil {
					return ""
				}
				return langue
			}
		}
	}

	return ""
}
