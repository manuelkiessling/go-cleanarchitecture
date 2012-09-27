package infrastructure

import (
	"fmt"
)

type Logger struct{}

func (logger Logger) Log(message string) error {
	fmt.Println("Log message: " + message)
	return nil
}
