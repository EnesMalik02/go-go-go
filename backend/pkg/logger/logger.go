package logger

import (
	"log"
	"os"
)

// Logger is a simple wrapper around the standard log package
type Logger struct {
	*log.Logger
}

// New creates a new Logger instance
func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[APP] ", log.LstdFlags),
	}
}

// Info logs informational messages
func (l *Logger) Info(v ...interface{}) {
	l.Println(v...)
}

// Error logs error messages
func (l *Logger) Error(v ...interface{}) {
	l.Println("ERROR:", v)
}
