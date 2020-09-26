package golog

import (
	"fmt"
	"log"
	"os"

	"github.com/recallsong/go-utils/logs"
)

// Logger .
type Logger struct {
	name   string
	prefix string
	log    *log.Logger
}

// New .
func New() *Logger {
	return &Logger{
		log: log.New(os.Stderr, "", log.LstdFlags),
	}
}

// Sub .
func (l *Logger) Sub(name string) logs.Logger {
	if len(l.name) <= 0 {
		return &Logger{
			name:   name,
			prefix: "[" + name + "]",
			log:    l.log,
		}
	}
	std := &Logger{
		name: l.name + "." + name,
		log:  l.log,
	}
	std.prefix = "[" + std.name + "]"
	return std
}

// Debug .
func (l *Logger) Debug(args ...interface{}) {
	l.log.Print(append([]interface{}{"[debug]", l.prefix}, args...)...)
}

// Info .
func (l *Logger) Info(args ...interface{}) {
	l.log.Println(append([]interface{}{"[info]", l.prefix}, args...)...)
}

// Warn .
func (l *Logger) Warn(args ...interface{}) {
	l.log.Println(append([]interface{}{"[warn]", l.prefix}, args...)...)
}

// Error .
func (l *Logger) Error(args ...interface{}) {
	l.log.Println(append([]interface{}{"[error]", l.prefix}, args...)...)
}

// Panic .
func (l *Logger) Panic(args ...interface{}) {
	l.log.Panic(append([]interface{}{"[panic]", l.prefix}, args...)...)
}

// Fatal .
func (l *Logger) Fatal(args ...interface{}) {
	l.log.Fatal(append([]interface{}{"[fatal]", l.prefix}, args...))
	os.Exit(-1)
}

// Debugf .
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.log.Println("[debug]", l.prefix, fmt.Sprintf(template, args...))
}

// Infof .
func (l *Logger) Infof(template string, args ...interface{}) {
	l.log.Println("[info]", l.prefix, fmt.Sprintf(template, args...))
}

// Warnf .
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.log.Println("[warn]", l.prefix, fmt.Sprintf(template, args...))
}

// Errorf .
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.log.Println("[error]", l.prefix, fmt.Sprintf(template, args...))
}

// Panicf .
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.log.Panicf("[panic]", l.prefix, fmt.Sprintf(template, args...))
}

// Fatalf .
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.log.Fatal("[fatal]", l.prefix, fmt.Sprintf(template, args...))
	os.Exit(-1)
}
