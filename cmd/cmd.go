package cmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	log "juby/logging"
	"juby/read"
	"os"
	"strings"
)

var (
	FilePath    string
	FileContent string
)

func findExt() {
	extensionArr := strings.Split(FilePath, ".")
	extension := "." + extensionArr[len(extensionArr)-1]
	log.Debug(fmt.Sprintf("File extension is %s", extension))
	result, err := read.GetExtensionType(extension)
	log.Debug(fmt.Sprint(err))
	if result.Name != "" {
		log.Info("Extension language: " + result.Name)
		log.Info(fmt.Sprintf("It is a %s language", result.Type))
		log.Info(fmt.Sprintf("All extensions of the programming language are: %s", strings.Join(result.Extensions, " ")))
	}
}

func Init() {
	flag.BoolVar(&log.DebugEnable, "debug", false, "enable detailed logging")
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Error("You did not input the file path right")
		os.Exit(0)
	}
	FilePath = flag.Args()[0]
	byteContent, err := ioutil.ReadFile(FilePath)
	if err != nil {
		log.Error(fmt.Sprint(err))
		os.Exit(0)
	}
	FileContent = string(byteContent)
	findExt()
}
