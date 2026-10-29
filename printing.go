package llg

import (
	"fmt"
	"os"
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
		fmt.Fprintln(os.Stderr, "[DEBUG] ")
		fmt.Fprintln(os.Stderr, v...)
	}
}
func PrintError(v ...any) {
	fmt.Fprint(os.Stderr, "[ERROR] ")
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}
