package printit

import (
	"fmt"
)

var (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

func Info(text string) {
	fmt.Println(text, string(colorReset))
}

func Success(text string) {
	fmt.Println(string(colorGreen), "✓", text, string(colorReset))
}

func Error(text string) {
	fmt.Println(string(colorRed), text, string(colorReset))
}
