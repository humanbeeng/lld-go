package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	logger := GetNewLogger(INFO)

	infLog, _ := NewFileAppender("info.log")
	errLog, _ := NewFileAppender("error.log")

	consoleAppender := ConsoleAppender{}
	logger.RegisterAppenders(INFO, &infLog, &consoleAppender)
	logger.RegisterAppenders(ERROR, &errLog, &consoleAppender)

	logger.Info("Hello this is an info messsage")
	logger.Error("Something went wrong while fetching movie details")
}
