package logger

import (
	"sync"
	"time"
)

type Logger struct {
	handler     LogHandler
	setLogLevel LogLevel
}

func newLogger(handler LogHandler, level LogLevel) *Logger {
	return &Logger{
		handler:     handler,
		setLogLevel: level,
	}
}

func GetNewLogger(level LogLevel) *Logger {
	loggerInit.Do(func() {
		logger = newLogger(*GetLogHandler(level), level)
	})

	return logger
}

func (l *Logger) RegisterAppenders(level LogLevel, appenders ...Appender) error {
	for _, a := range appenders {
		l.handler.AddAppender(level, a)
	}
	return nil
}

func (l *Logger) Debug(msg string) {
	record := Record{
		Message:   msg,
		Level:     DEBUG,
		Timestamp: time.Now(),
	}
	l.handler.Handle(record)
}

func (l *Logger) Info(msg string) {
	record := Record{
		Message:   msg,
		Level:     INFO,
		Timestamp: time.Now(),
	}
	l.handler.Handle(record)
}

func (l *Logger) Warn(msg string) {
	record := Record{
		Message:   msg,
		Level:     WARN,
		Timestamp: time.Now(),
	}
	l.handler.Handle(record)
}

func (l *Logger) Error(msg string) {
	record := Record{
		Message:   msg,
		Level:     ERROR,
		Timestamp: time.Now(),
	}
	l.handler.Handle(record)
}

func (l *Logger) Fatal(msg string) {
	record := Record{
		Message:   msg,
		Level:     FATAL,
		Timestamp: time.Now(),
	}
	l.handler.Handle(record)
}

var (
	logger     *Logger
	loggerInit sync.Once
)
