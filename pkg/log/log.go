package log

var (
	Debug Logger = NewDebugLogger()
	Info  Logger = NewInfoLogger()
	Warn  Logger = NewWarnLogger()
	Error Logger = NewErrorLogger()
)

type Logger interface {
	Print(v ...any)
	Println(v ...any)
	Printf(format string, v ...any)
}
