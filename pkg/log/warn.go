package log

import (
	"log"
	"os"
)

type WarnLogger struct {
	log.Logger
}

func NewWarnLogger() *WarnLogger {
	return &WarnLogger{
		Logger: *log.New(
			os.Stderr,
			"\033[1;33m[ WARN]\033[0m ",
			log.LstdFlags|log.Lshortfile,
		),
	}
}