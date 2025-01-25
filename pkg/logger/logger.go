package logger

import (
	"fmt"
	"os"
)

var Info = "\033[32m✔\033[0m"
var Error = "\033[31m✘\033[0m"
var Warning = "\033[33m⚠\033[0m"

func INFO(format string, a ...any) {
	fmt.Printf("[%s] %s\n", Info, fmt.Sprintf(format, a...))
}

func WARNING(format string, a ...any) {
	fmt.Printf("[%s] %s\n", Warning, fmt.Sprintf(format, a...))
}

func ERROR(format string, a ...any) {
	fmt.Printf("[%s] %s\n", Error, fmt.Sprintf(format, a...))
	os.Exit(1)
}
