package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type ErrorLogger struct {
	log.Logger
}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{
		Logger: *log.New(
			os.Stderr,
			"\033[1;31m[ERROR]\033[0m ",
			log.LstdFlags|log.Lshortfile,
		),
	}
}

func (l *ErrorLogger) Print(v ...any) {
	_, file, line, _ := runtime.Caller(1)
	fileSplit := strings.Split(file, "/")
	prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ": "
	l.Logger.Print(prefix + "\033[31m" + fmt.Sprint(v...) + "\033[0m")
}

func (l *ErrorLogger) Println(v ...any) {
	_, file, line, _ := runtime.Caller(1)
	fileSplit := strings.Split(file, "/")
	prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ": "
	l.Logger.Println(prefix + "\033[31m" + fmt.Sprint(v...) + "\033[0m")
}

func (l *ErrorLogger) Printf(format string, v ...any) {
	_, file, line, _ := runtime.Caller(1)
	fileSplit := strings.Split(file, "/")
	prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ":"
	l.Logger.Printf(
		"%s %s",
		prefix,
		fmt.Sprintf("\033[31m"+format+"\033[0m", v...),
	)
}
