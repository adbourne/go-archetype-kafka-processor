package config

import (
	"log"
	"time"
)

// Abstraction for logging
// We can use this interface throughout the application and swap out the underlying logging
// libraries at will
type Logger interface {
	Error(msg string)

	Warn(msg string)

	Info(msg string)

	Debug(msg string)

	Trace(msg string)
}

// Create a new Logger
// TODO: How to make this present everywhere, but separate from the AppContext? Perhaps this is good enough
func NewLogger() Logger {
	return &SystemOutLogger{}
}

// Logs everything to system out
type SystemOutLogger struct {
}

func (_ *SystemOutLogger) Error(msg string) {
	logLine("ERROR", msg)
}

func (_ *SystemOutLogger) Warn(msg string) {
	logLine("WARN", msg)
}

func (_ *SystemOutLogger) Info(msg string) {
	logLine("INFO", msg)
}

func (_ *SystemOutLogger) Debug(msg string) {
	logLine("DEBUG", msg)
}

func (_ *SystemOutLogger) Trace(msg string) {
	logLine("TRACE", msg)
}

func logLine(level string, msg string) {
	log.Printf("%s %s: %s", getTime(), level, msg)
}

func getTime() string {
	return time.Now().Format(time.RFC3339)
}
