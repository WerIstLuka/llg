package llg

import (
	"fmt"
	"os"
	"strings"
)

var isDebug bool = false

func CheckDebug() {
	if os.Getenv("DEBUG") == "true" {
		isDebug = true
	}
}

func GetDebugPtr() *bool {
	return &isDebug
}

func PrintDebug(v ...any) {
	if isDebug {
		fmt.Fprintln(os.Stderr, "[DEBUG]", strings.Trim(fmt.Sprint(v...), "[]"))
	}
}
func PrintError(v ...any) {
	fmt.Fprintln(os.Stderr, "[ERROR]", strings.Trim(fmt.Sprint(v...), "[]"))
	os.Exit(1)
}
