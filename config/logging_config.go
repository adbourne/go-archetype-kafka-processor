package config

import (
	"log"
	"time"
)

// Logger is an abstraction for logging
// We can use this interface throughout the application and swap out the underlying logging
// libraries at will
type Logger interface {
	Error(msg string)

	Warn(msg string)

	Info(msg string)

	Debug(msg string)

	Trace(msg string)
}

// NewLogger creates a new Logger
// TODO: How to make this present everywhere, but separate from the AppContext? Perhaps this is good enough
func NewLogger() Logger {
	return &SystemOutLogger{}
}

// SystemOutLogger logs everything to system out
type SystemOutLogger struct {
}

// Error is an ERROR level log message
func (l *SystemOutLogger) Error(msg string) {
	logLine("ERROR", msg)
}

// Warn is an WARN level log message
func (l *SystemOutLogger) Warn(msg string) {
	logLine("WARN", msg)
}

// Info is an INFO level log message
func (l *SystemOutLogger) Info(msg string) {
	logLine("INFO", msg)
}

// Debug is an DEBUG level log message
func (l *SystemOutLogger) Debug(msg string) {
	logLine("DEBUG", msg)
}

// Trace is an TRACE level log message
func (l *SystemOutLogger) Trace(msg string) {
	logLine("TRACE", msg)
}

// logLine logs a line in a standard format using the provided level and message
func logLine(level string, msg string) {
	log.Printf("%s %s: %s", getTime(), level, msg)
}

// getTime gets the time in RFC3339 format
func getTime() string {
	return time.Now().Format(time.RFC3339)
}
