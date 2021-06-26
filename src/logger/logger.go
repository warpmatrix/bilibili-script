package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
	fatal *log.Logger
)

func init() {
	info = log.New(os.Stdout, "[INFO]", log.Ldate)
	warn = log.New(os.Stderr, "[WARN]", log.Ldate)
	err = log.New(os.Stderr, "[ERROR]", log.LstdFlags|log.Lshortfile)
	fatal = log.New(io.MultiWriter(os.Stderr), "[FATAL]", log.LstdFlags|log.Lshortfile)
}

func Info(v ...interface{}) {
	info.Output(2, fmt.Sprintln(v...))
}

func Warning(v ...interface{}) {
	warn.Output(2, fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	err.Output(2, fmt.Sprintln(v...))
}

func Fatal(v ...interface{}) {
	s := fmt.Sprintln(v...)
	fatal.Output(2, s)
	panic(s)
}
