package log

import (
	"log"
	"os"
)

type InfoLogger struct {
	log.Logger
}

func NewInfoLogger() *InfoLogger {
	return &InfoLogger{
		Logger: *log.New(
			os.Stderr,
			"\033[1;36m[ INFO]\033[0m ",
			log.LstdFlags|log.Lshortfile,
		),
	}
}
