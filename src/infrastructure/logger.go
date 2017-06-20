package infrastructure

import (
	"log"
)

type Logger struct{}

func (logger Logger) Log(args ...interface{}) {
	log.Println(args...)
}
