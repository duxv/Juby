package utils

import (
	"flag"
	"fmt"
	"io/ioutil"
	"juby/read"
	"os"
	"strings"
)

var (
	DebugEnable bool
	FilePath    string
	FileContent string
)

func findExt() {
	extensionArr := strings.Split(FilePath, ".")
	extension := extensionArr[len(extensionArr)-1]
	extension = "." + extension
	Debug(fmt.Sprintf("File extension is %s", extension))
	result, err := read.GetExtensionType(extension)
	Debug(fmt.Sprint(err))
	if result.Name != "" {
		Info("Found file extension name: " + result.Name)
		Info("The extension's programming language type: " + result.Type)
		Info(fmt.Sprintf("All extensions of the programming language are: %v", result.Extensions))
	}
}

func Init() {
	flag.BoolVar(&DebugEnable, "debug", false, "enable detailed logging")
	flag.Parse()
	if len(flag.Args()) != 1 {
		Error("You did not input the file path right")
		os.Exit(0)
	}
	FilePath = flag.Args()[0]
	byteContent, err := ioutil.ReadFile(FilePath)
	if err != nil {
		Error(fmt.Sprint(err))
		os.Exit(0)
	}
	FileContent = string(byteContent)
	findExt()
}
