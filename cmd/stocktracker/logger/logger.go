package logger

import (
	"fmt"
	"strings"

	"github.com/golobby/config"
)

// Logger provides basic logging functions to write to the console
type Logger struct {
	config *config.Config
}

// ProvideLogger provides a new instance for wire
func ProvideLogger(config *config.Config) *Logger {
	return &Logger{config}
}

var (
	traceLevel = 1
	debugLevel = 2
	infoLevel  = 3
	warnLevel  = 4
	fatalLevel = 5
)

var (
	trace = magenta
	debug = green
	info  = teal
	warn  = yellow
	fatal = red
)

var (
	black   = color("\033[1;30m%s\033[0m")
	red     = color("\033[1;31m%s\033[0m")
	green   = color("\033[1;32m%s\033[0m")
	yellow  = color("\033[1;33m%s\033[0m")
	purple  = color("\033[1;34m%s\033[0m")
	magenta = color("\033[1;35m%s\033[0m")
	teal    = color("\033[1;36m%s\033[0m")
	white   = color("\033[1;37m%s\033[0m")
)

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func (l *Logger) getLogLevel() int {
	logLevel, err := l.config.GetString("logLevel")

	if err != nil {
		panic(err)
	}

	switch logLevel {
	case "trace":
		fallthrough
	case "Trace":
		return traceLevel
	case "debug":
		fallthrough
	case "Debug":
		return debugLevel
	case "info":
		fallthrough
	case "Info":
		return infoLevel
	case "warning":
		fallthrough
	case "Warning":
		return warnLevel
	case "fatal":
		fallthrough
	case "Fatal":
		return fatalLevel
	default:
		return 0
	}
}

func (l *Logger) canPrint(level int) bool {
	return l.getLogLevel() <= level
}

// LogTrace writes to the console
func (l *Logger) LogTrace(message ...string) {
	if l.canPrint(traceLevel) {
		message = append([]string{"[TRACE]"}, message...)
		l.Log(trace, message...)
	}
}

// LogDebug writes to the console
func (l *Logger) LogDebug(message ...string) {
	if l.canPrint(debugLevel) {
		message = append([]string{"[DEBUG]"}, message...)
		l.Log(debug, message...)
	}
}

// LogInfo writes to the console
func (l *Logger) LogInfo(message ...string) {
	if l.canPrint(infoLevel) {
		message = append([]string{"[INFO]"}, message...)
		l.Log(info, message...)
	}
}

// LogWarning writes to the console
func (l *Logger) LogWarning(message ...string) {
	if l.canPrint(warnLevel) {
		message = append([]string{"[WARN]"}, message...)
		l.Log(warn, message...)
	}
}

// LogFatal writes to the console
func (l *Logger) LogFatal(message ...string) {
	if l.canPrint(fatalLevel) {
		message = append([]string{"[FATAL]"}, message...)
		l.Log(fatal, message...)
	}
}

// Log writes an array of string to the console
func (l *Logger) Log(logLevel func(...interface{}) string, message ...string) {
	if logLevel != nil {
		fmt.Println(logLevel(strings.Join(message, " ")))
	} else {
		fmt.Println(strings.Join(message, " "))
	}
}
