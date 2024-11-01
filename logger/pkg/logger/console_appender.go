package logger

import "fmt"

type ConsoleAppender struct{}

func (a *ConsoleAppender) Append(msg string) {
	fmt.Println(msg)
}
