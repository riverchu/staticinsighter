package analyzer

import (
	"fmt"
	"strings"
	"time"
)

type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel

	DateTime = "2006-01-02 15:04:05"
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "Trace"
	case DebugLevel:
		return "Debug"
	case InfoLevel:
		return "Info"
	case WarnLevel:
		return "Warn"
	case ErrorLevel:
		return "Error"
	case FatalLevel:
		return "Fatal"
	default:
		return ""
	}
}

func (l Level) Color() int {
	switch l {
	case TraceLevel:
		return 45
	case DebugLevel:
		return 39
	case InfoLevel:
		return 33
	case WarnLevel:
		return 148
	case ErrorLevel:
		return 161
	case FatalLevel:
		return 160
	default:
		return 0
	}
}

// Logger logger interface
type Logger interface {
	SetLevel(level Level)

	Trace(format string, v ...any)
	Debug(format string, v ...any)
	Info(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
	Fatal(format string, v ...any)
}

func defaultLogger() Logger { return &BuiltinLogger{level: InfoLevel} }

type BuiltinLogger struct {
	level Level
}

func (l *BuiltinLogger) SetLevel(level Level)        { l.level = level }
func (l *BuiltinLogger) allowLevel(level Level) bool { return level >= l.level }

func (l *BuiltinLogger) Trace(format string, v ...any) {
	if !l.allowLevel(TraceLevel) {
		return
	}
	fmt.Printf(l.format(TraceLevel, format), v...)
}
func (l *BuiltinLogger) Debug(format string, v ...any) {
	if !l.allowLevel(DebugLevel) {
		return
	}
	fmt.Printf(l.format(DebugLevel, format), v...)
}
func (l *BuiltinLogger) Info(format string, v ...any) {
	if !l.allowLevel(InfoLevel) {
		return
	}
	fmt.Printf(l.format(InfoLevel, format), v...)
}
func (l *BuiltinLogger) Warn(format string, v ...any) {
	if !l.allowLevel(WarnLevel) {
		return
	}
	fmt.Printf(l.format(WarnLevel, format), v...)
}
func (l *BuiltinLogger) Error(format string, v ...any) {
	if !l.allowLevel(ErrorLevel) {
		return
	}
	fmt.Printf(l.format(ErrorLevel, format), v...)
}
func (l *BuiltinLogger) Fatal(format string, v ...any) {
	if !l.allowLevel(FatalLevel) {
		return
	}
	fmt.Printf(l.format(FatalLevel, format), v...)
}

func (l *BuiltinLogger) format(level Level, format string) string {
	return fmt.Sprintf("\033[38;5;%dm%s [%s]%s\033[0m\n", level.Color(), time.Now().Format(DateTime), strings.ToUpper(level.String()), format)
}
