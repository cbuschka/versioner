package cli

import (
	"fmt"
	"os"
)

// Die prints a message and exists the process with exitCode.
func Die(exitCode int, format string, a ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
	os.Exit(exitCode)
}

// Print prints a message to console.
func Print(format string, a ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
}

// Debug prints a message to console when verbose mode is chosen.
func Debug(format string, a ...interface{}) {
}
