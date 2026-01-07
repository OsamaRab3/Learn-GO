package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Level int

const (
	INFO Level = iota
	DEBUG
	WARN
	ERROR
)

type Logger struct {
	level Level
	out   io.Writer
}

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func (l *Logger) log(level Level, message string) {
	if level < l.level {
		return
	}
	fmt.Fprintf(l.out, "%s [%s] %s\n", time.Now().Format(time.RFC3339), level.String(), message)

}
func (l *Logger) Info(message string) {
	l.log(INFO, message)
}

func (l *Logger) Error(message string) {
	l.log(ERROR, message)
}

func (l *Logger) Warn(message string) {
	l.log(WARN, message)
}

func (l *Logger) Debug(message string) {
	l.log(DEBUG, message)
}

func main() {
	logger := &Logger{
		level: INFO,
		out:   os.Stdout,
	}
	logger.Info("server started")
	logger.Debug("this will not show")
	logger.Warn("this is a warn")
	logger.Error("this is a Error")

}
