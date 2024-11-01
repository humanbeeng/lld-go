package logger

type Appender interface {
	Append(msg string)
}
