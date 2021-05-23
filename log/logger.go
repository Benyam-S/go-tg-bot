package log

import (
	"fmt"
	"os"
	"time"
)

type ILogger interface {
	Log(stmt, stmtType, logFile string)
	LogToParent(stmt, stmtType string)
	LogFileError(stmt, logFile string)
}

// LogBug is a type that defines a log bug that contains the log files location
type LogBug struct {
	ServerLogFile string
	BotLogFile    string
	Logger        ILogger
}

// Logger is a type that defines a logger type
type Logger struct{}

// Log is a method that will log the given statement to the selected log file
func (l *Logger) Log(stmt, stmtType, logFile string) {

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		defer file.Close()

		if stmtType == "w" {
			stmtType = "Warning:"

		} else if stmtType == "e" {
			stmtType = "Error:"

		}

		stmt = fmt.Sprintf("%s %s  At %s", stmtType, stmt, time.Now())

		fmt.Fprintln(file, stmt)

	}
}

// LogToParent is a method that will log the given statement to the program starter
func (l *Logger) LogToParent(stmt, stmtType string) {

	if stmtType == "w" {
		stmtType = "Warning:"

	} else if stmtType == "e" {
		stmtType = "Error:"

	} else {
		stmtType = "Success:"

	}

	stmt = fmt.Sprintf("%s %s", stmtType, stmt)
	fmt.Println(stmt)
}

// LogFileError is a method that will log the given statement as an error to the selected log file
func (l *Logger) LogFileError(stmt, logFile string) {

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		defer file.Close()

		stmt = fmt.Sprintf("Error: %s  At %s", stmt, time.Now())

		fmt.Fprintln(file, stmt)

	}
}
