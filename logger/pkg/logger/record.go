package logger

import (
	"fmt"
	"strings"
	"time"
)

type Record struct {
	Level      LogLevel
	Message    string
	Attributes RecordAttributes
	Timestamp  time.Time
}

type RecordAttributes map[string]any

func (a RecordAttributes) String() string {
	var str string
	var attrs []string
	for k, v := range a {
		attrs = append(attrs, fmt.Sprintf("%v=%v", k, v))
	}

	str = strings.Join(attrs, " ")
	return str
}

func (r *Record) String() string {
	return fmt.Sprintf("%v [%v] %v %v\n", r.Timestamp.Format(time.RFC3339), r.Level, r.Message, r.Attributes.String())
}
