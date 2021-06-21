package utils

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// Print the detailed content
func Debug(content string) {
	if DebugEnable {
		fmt.Printf("[%v] %s\n", aurora.Blue("DBG"), content)
	}
}

func Info(content string) {
	fmt.Printf("[%v] %s\n", aurora.Green("INF"), content)
}

func Warn(content string) {
	fmt.Printf("[%v] %s\n", aurora.Yellow("WRN"), content)
}

func Error(content string) {
	fmt.Printf("[%v] %s\n", aurora.Red("ERR"), content)
}
