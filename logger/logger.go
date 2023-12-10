package logger

import (
	"log"
	"os"
)

// Logger is a simple logger that wraps the standard log package.
type Logger struct {
	*log.Logger
}

// NewLogger creates a new Logger instance.
func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "GraphQL_GoLang_CRUD_Application", log.LstdFlags),
	}
}
