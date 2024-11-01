package logger

import (
	"sync"
)

type LogHandler struct {
	level     LogLevel
	observers map[LogLevel][]Appender
}

func GetLogHandler(level LogLevel) *LogHandler {
	logHandlerInit.Do(func() {
		handler := newLogHandler(level)
		logHandler = &handler
	})

	return logHandler
}

func newLogHandler(level LogLevel) LogHandler {
	return LogHandler{
		level:     level,
		observers: make(map[LogLevel][]Appender),
	}
}

func (h *LogHandler) AddAppender(level LogLevel, appender Appender) error {
	appenders, ok := h.observers[level]
	if !ok {
		h.observers[level] = []Appender{appender}
		return nil
	}

	appenders = append(appenders, appender)
	h.observers[level] = appenders
	return nil
}

func (h *LogHandler) Handle(record Record) {
	if record.Level < LogLevel(h.level) {
		return
	}

	appenders := h.observers[record.Level]

	for _, a := range appenders {
		a.Append(record.String())
	}
}

var (
	logHandler     *LogHandler
	logHandlerInit sync.Once
)
