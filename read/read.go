package read

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Extension struct {
	Name       string
	Type       string
	Extensions []string
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
