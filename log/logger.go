package log

import (
    "fmt"
    "io"
)

// Logger TODO
type Logger struct {
    w io.Writer
}

// New TODO
func New(writer io.Writer) *Logger {
    return &Logger{
        w: writer,
    }
}

//
func (l *Logger) Error(msg string, err error) {
    _, _ = l.w.Write([]byte(fmt.Sprintf("%q: %v", msg, err)))
}
