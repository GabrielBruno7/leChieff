package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		error:   log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (logger *Logger) Debug(value ...interface{}) {
	logger.debug.Println(value...)
}

func (logger *Logger) Info(value ...interface{}) {
	logger.info.Println(value...)
}

func (logger *Logger) Warning(value ...interface{}) {
	logger.warning.Println(value...)
}

func (logger *Logger) Error(value ...interface{}) {
	logger.error.Println(value...)
}

func (logger *Logger) DebugFormatted(format string, value ...interface{}) {
	logger.debug.Printf(format, value...)
}

func (logger *Logger) InfoFormatted(format string, value ...interface{}) {
	logger.info.Printf(format, value...)
}

func (logger *Logger) WarningFormatted(format string, value ...interface{}) {
	logger.warning.Printf(format, value...)
}

func (logger *Logger) ErrorFormatted(format string, value ...interface{}) {
	logger.error.Printf(format, value...)
}
