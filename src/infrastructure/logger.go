package infrastructure

import (
	"fmt"
)

type Logger struct {}

function (logger Logger) Log(message string) error {
	fmt.Println("Log message: " + message)
}
