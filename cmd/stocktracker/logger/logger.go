package logger

import (
	"fmt"

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
	errorLevel = 5
)

var (
	trace = magenta
	debug = green
	info  = teal
	warn  = yellow
	error = red
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

func (logger *Logger) getLogLevel() int {
	logLevel, err := logger.config.GetString("logLevel")

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
	case "error":
		fallthrough
	case "Error":
		return errorLevel
	default:
		return 0
	}
}

func (logger *Logger) canPrint(level int) bool {
	return logger.getLogLevel() <= level
}

// LogTrace writes to the console
func (logger *Logger) LogTrace(message string) {
	if logger.canPrint(traceLevel) {
		fmt.Println(trace("[TRACE] ", message))
	}
}

// LogDebug writes to the console
func (logger *Logger) LogDebug(message string) {
	if logger.canPrint(debugLevel) {
		fmt.Println(debug("[DEBUG] ", message))
	}
}

// LogInfo writes to the console
func (logger *Logger) LogInfo(message string) {
	if logger.canPrint(infoLevel) {
		fmt.Println(info("[INFO] ", message))
	}
}

// LogWarning writes to the console
func (logger *Logger) LogWarning(message string) {
	if logger.canPrint(warnLevel) {
		fmt.Println(warn("[WARN] ", message))
	}
}

// LogError writes to the console
func (logger *Logger) LogError(message string) {
	if logger.canPrint(errorLevel) {
		fmt.Println(error("[ERROR] ", message))
	}
}
