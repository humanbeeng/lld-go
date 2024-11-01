package logger_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/humanbeeng/lld-go/logger/pkg/logger"
)

func TestRecordString(t *testing.T) {
	r := logger.Record{
		Level:      logger.INFO,
		Timestamp:  time.Now(),
		Message:    "User created",
		Attributes: map[string]any{"name": "nithin", "age": 1},
	}

	msg := r.String()
	fmt.Print(msg)

}
