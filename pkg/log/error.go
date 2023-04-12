package log

import (
	"log"
	"os"
)

type ErrorLogger struct {
	log.Logger
}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{
		Logger: *log.New(
			os.Stderr,
			"\033[1;3;37m[ERROR]\033[0m ",
			log.LstdFlags|log.Lshortfile,
		),
	}
}
