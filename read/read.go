package read

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"juby/logging"
	"strings"
)

type Language struct {
	Name     string
	Found    int
	NotFound int
	Keywords []string `json:"keys"`
}

type Extension struct {
	Name       string
	Type       string
	Extensions []string
}

func Contains(key string, substrings []string) bool {
	for _, r := range substrings {
		if strings.Contains(key, r) {
			return true
		}
	}
	return false
}

func ContainsArr(key string, stringz []string) bool {
	for _, str := range stringz {
		if str == key {
			return true
		}
	}
	return false
}
func FindLanguage(content string) (map[string]int, error) {
	var langs []Language
	// Count of the keys checked in the content
	scannedKeysCount := 0
	var scannedKeys []string
	var inString string
	var quotes = []string{"'", "\"", `"""`, "`"}
	keywords, err := Readfile("./assets/keywords.json")
	keywords = strings.NewReplacer("\n", " ", ";", " ").Replace(keywords)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(keywords), &langs)
	for _, key := range strings.Split(content, " ") {
		if inString == "" && Contains(key, quotes) {
			inString = string(key[0])
			continue
		}
		if inString != "" && strings.Contains(key, inString) {
			inString = ""
			continue
		}
		if inString != "" {
			continue
		}
		for j := 0; j < len(langs); j++ {
			for i := 0; i < len(langs[j].Keywords); i++ {
				if key == "" || ContainsArr(key, scannedKeys) {
					continue
				}
				if strings.Contains(key, langs[j].Keywords[i]) {
					langs[j].Keywords[i] = key
					langs[j].Found += 1
					scannedKeys = append(scannedKeys, langs[j].Keywords[i])
					logging.Debug(fmt.Sprintf("Found key %s from programming language %s", langs[j].Keywords[i], langs[j].Name))
				}
			}
		}
		scannedKeysCount += 1
	}
	// Make the result map
	result := make(map[string]int)
	for _, lang := range langs {
		result[lang.Name] = lang.Found
	}
	return result, nil
}

func Readfile(filepath string) (res string, er error) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func GetExtensionType(extension string) (ext Extension, er error) {
	var extensions []Extension
	var result Extension
	fileExtensions, err := Readfile("./assets/langs.json")
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal([]byte(fileExtensions), &extensions); err != nil {
		return result, err
	}
	for l, j := range extensions {
		for _, i := range j.Extensions {
			if i == extension {
				result = extensions[l]
				return result, nil
			}
		}
	}
	return result, errors.New("file extension information not found")
}
