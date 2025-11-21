package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	// ANSI color codes
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

var (
	// Create a custom logger without any flags (no timestamp, no prefix)
	colorLogger = log.New(os.Stderr, "", 0)
)

// Println prints a colored log message with timestamp in cyan and message in green
func Println(v ...interface{}) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	message := fmt.Sprintln(v...)
	// Remove the trailing newline from Sprintln
	if len(message) > 0 && message[len(message)-1] == '\n' {
		message = message[:len(message)-1]
	}
	colorLogger.Printf("%s%s%s %s%s%s", colorCyan, timestamp, colorReset, colorGreen, message, colorReset)
}

// Fatalf prints a colored fatal error message with timestamp in cyan and message in yellow, then exits
func Fatalf(format string, v ...interface{}) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	message := fmt.Sprintf(format, v...)
	colorLogger.Fatalf("%s%s%s %s%s%s", colorCyan, timestamp, colorReset, colorYellow, message, colorReset)
}

// SetOutput sets the output destination for the logger
func SetOutput(w io.Writer) {
	colorLogger.SetOutput(w)
}
