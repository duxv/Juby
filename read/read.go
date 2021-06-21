package read

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

type Language struct {
	Name     string
	Score    int
	Keywords []string
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

func FindLanguage(content string) error {
	var langs []Language
	var inString string
	var quotes = []string{"'", "\""}
	keywords, err := Readfile("./assets/keywords.json")
	if err != nil {
		return err
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
		// TODO: check all programming languages and do an algorithm to decide which programming language is
	}
	return nil
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
