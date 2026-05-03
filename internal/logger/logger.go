package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
)

var noColor bool

func init() {
	// Disable color if not a TTY (e.g. piped or systemd)
	if os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		noColor = true
	}
}

func timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func colorize(color, text string) string {
	if noColor {
		return text
	}
	return color + text + colorReset
}

func Info(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s %s\n",
		colorize(colorGray, timestamp()),
		colorize(colorGreen, "[INFO]"),
		msg,
	)
}

func Error(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s %s %s\n",
		colorize(colorGray, timestamp()),
		colorize(colorRed, "[ERROR]"),
		msg,
	)
}

func Warn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s %s\n",
		colorize(colorGray, timestamp()),
		colorize(colorYellow, "[WARN]"),
		msg,
	)
}

func Request(method, host, path string, status int, duration time.Duration) {
	statusColor := colorGreen
	if status >= 400 && status < 500 {
		statusColor = colorYellow
	} else if status >= 500 {
		statusColor = colorRed
	}

	fmt.Printf("%s %s %s %s%s → %s (%s)\n",
		colorize(colorGray, timestamp()),
		colorize(colorCyan, "[REQ]"),
		colorize(colorCyan, method),
		host,
		path,
		colorize(statusColor, fmt.Sprintf("%d", status)),
		duration.Round(time.Millisecond),
	)
}

func Fatal(format string, args ...any) {
	Error(format, args...)
	os.Exit(1)
}
