package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
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

func (l *InfoLogger) Print(v ...any) {
	_, file, line, _ := runtime.Caller(1)
	fileSplit := strings.Split(file, "/")
	prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ": "
	l.Logger.Print(prefix + "\033[36m" + fmt.Sprint(v...) + "\033[0m")
}

func (l *InfoLogger) Println(v ...any) {
	_, file, line, _ := runtime.Caller(1)
	fileSplit := strings.Split(file, "/")
	prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ": "
	l.Logger.Println(prefix + "\033[36m" + fmt.Sprint(v...) + "\033[0m")
}

func (l *InfoLogger) Printf(format string, v ...any) {
	_, file, line, _ := runtime.Caller(1)
	fileSplit := strings.Split(file, "/")
	prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ":"
	l.Logger.Printf(
		"%s %s",
		prefix,
		fmt.Sprintf("\033[36m"+format+"\033[0m", v...),
	)
}
