package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type DebugLogger struct {
	log.Logger

	Enabled bool
}

func NewDebugLogger() *DebugLogger {
	return &DebugLogger{
		Logger:  *log.New(os.Stderr, "\033[1;3;37m[DEBUG]\033[0m ", log.LstdFlags),
		Enabled: os.Getenv("DEBUG") == "true" || os.Getenv("DEBUG") == "yes",
	}
}

func (l *DebugLogger) Print(v ...any) {
	if l.Enabled {
		_, file, line, _ := runtime.Caller(1)
		fileSplit := strings.Split(file, "/")
		prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ": "
		l.Logger.Print(prefix + "\033[3;37m" + fmt.Sprint(v...) + "\033[0m")
	}
}

func (l *DebugLogger) Println(v ...any) {
	if l.Enabled {
		_, file, line, _ := runtime.Caller(1)
		fileSplit := strings.Split(file, "/")
		prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ": "
		l.Logger.Println(prefix + "\033[3;37m" + fmt.Sprint(v...) + "\033[0m")
	}
}

func (l *DebugLogger) Printf(format string, v ...any) {
	if l.Enabled {
		_, file, line, _ := runtime.Caller(1)
		fileSplit := strings.Split(file, "/")
		prefix := fileSplit[len(fileSplit)-1] + ":" + strconv.Itoa(line) + ":"
		l.Logger.Printf(
			"%s %s",
			prefix,
			fmt.Sprintf("\033[3;37m"+format+"\033[0m", v...),
		)
	}
}
