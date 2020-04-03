package logger

import "fmt"

type Logger struct{}

func (logger *Logger) Log(message string) {
	fmt.Println(message)
}
