package logger

import (
	"io"
	"log"
)

type Level int

const (
	// Error is messages about unexpected conditions that prevents normal command execution
	Error Level = iota
	// Warning is messages about abnormal but not critical things
	Warning
	// Info level is common debugging infomation. Should contain messages for developers
	Info
	// Debug level is only for development, no debug logs in master
	Debug
)

const flags = log.LstdFlags

type Logger interface {
	Debug(v ...any)
	Info(v ...any)
	Warning(v ...any)
	Error(v ...any)
}

func New(stdout, stderr io.Writer, level Level) Logger {
	l := logger{
		level: level,
	}
	l.error = log.New(stderr, "ERROR: ", flags)
	if level >= Warning {
		l.warn = log.New(stderr, "WARNING: ", flags)
	}
	if level >= Debug {
		l.debug = log.New(stdout, "DEBUG: ", flags)
	}
	if level >= Info {
		l.info = log.New(stdout, "INFO: ", flags)
	}
	return &l
}

type logger struct {
	level                    Level
	debug, info, warn, error *log.Logger
}

func (l *logger) Debug(v ...any) {
	if l.debug != nil {
		l.debug.Println(v...)
	}
}

func (l *logger) Info(v ...any) {
	if l.info != nil {
		l.info.Println(v...)
	}
}

func (l *logger) Warning(v ...any) {
	if l.warn != nil {
		l.warn.Println(v...)
	}
}

func (l *logger) Error(v ...any) {
	if l.error != nil {
		l.error.Println(v...)
	}
}
