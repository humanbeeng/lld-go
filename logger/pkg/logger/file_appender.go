package logger

import (
	"bufio"
	"fmt"
	"os"
)

type FileAppender struct {
	file *os.File
}

func NewFileAppender(filepath string) (FileAppender, error) {
	logFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return FileAppender{}, err
	}

	return FileAppender{
		file: logFile,
	}, nil
}

func (a *FileAppender) Append(msg string) {
	w := bufio.NewWriter(a.file)
	defer w.Flush()
	_, err := w.WriteString(msg)
	if err != nil {
		fmt.Println("Cannot append to file", err)
	}
}
