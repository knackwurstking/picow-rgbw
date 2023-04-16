package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
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

func (l *WarnLogger) Print(v ...any) {
	_, file, line, _ := runtime.Caller(1)
	fileSplit := strings.Split(file, "/")
	prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ": "
	l.Logger.Print(prefix + "\033[33m" + fmt.Sprint(v...) + "\033[0m")
}

func (l *WarnLogger) Println(v ...any) {
	_, file, line, _ := runtime.Caller(1)
	fileSplit := strings.Split(file, "/")
	prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ": "
	l.Logger.Println(prefix + "\033[33m" + fmt.Sprint(v...) + "\033[0m")
}

func (l *WarnLogger) Printf(format string, v ...any) {
	_, file, line, _ := runtime.Caller(1)
	fileSplit := strings.Split(file, "/")
	prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ":"
	l.Logger.Printf(
		"%s %s",
		prefix,
		fmt.Sprintf("\033[33m"+format+"\033[0m", v...),
	)
}
