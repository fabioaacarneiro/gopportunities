package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
    debug *log.Logger
    info *log.Logger
    warning *log.Logger
    error *log.Logger
    writer io.Writer
}

func NewLogger(p string) *Logger {
    writer := io.Writer(os.Stdout)
    logger := log.New(writer, p, log.Ldate|log.Ltime)

    return &Logger{
        debug: log.New(writer, "DEBUG: ", logger.Flags()),
        info: log.New(writer, "INFO: ", logger.Flags()),
        warning: log.New(writer, "WARNING: ", logger.Flags()),
        error: log.New(writer, "ERROR: ", logger.Flags()),
        writer: writer,
    }
}

func (l *Logger) Debug(v ...any) {
    l.debug.Println(v ...)
}

func (l *Logger) Info(v ...any) {
    l.info.Println(v ...)
}

func (l *Logger) Warning(v ...any) {
    l.warning.Println(v ...)
}

func (l *Logger) Error(v ...any) {
    l.error.Println(v ...)
}

func (l *Logger) Debugf(format string, v ...any) {
    l.debug.Printf(format, v ...)
}

func (l *Logger) Infof(format string, v ...any) {
    l.info.Printf(format, v ...)
}

func (l *Logger) Warningf(format string, v ...any) {
    l.warning.Printf(format, v ...)
}

func (l *Logger) Errorf(format string, v ...any) {
    l.error.Printf(format, v ...)
}
