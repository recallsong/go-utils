package stdout

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/recallsong/go-utils/logs"
)

// TimeFormat .
const TimeFormat = "2006-01-02 15:04:05.000"

// Stdout .
type Stdout struct {
	name   string
	prefix string
}

// Sub .
func (l *Stdout) Sub(name string) logs.Logger {
	if len(l.name) <= 0 {
		return &Stdout{
			name:   name,
			prefix: "[" + name + "]",
		}
	}
	std := &Stdout{
		name: l.name + "." + name,
	}
	std.prefix = "[" + std.name + "]"
	return std
}

// Debug .
func (l *Stdout) Debug(args ...interface{}) {
	fmt.Println(append([]interface{}{time.Now().Format(TimeFormat), "[debug]", l.prefix}, args...)...)
}

// Info .
func (l *Stdout) Info(args ...interface{}) {
	fmt.Println(append([]interface{}{time.Now().Format(TimeFormat), "[info]", l.prefix}, args...)...)
}

// Warn .
func (l *Stdout) Warn(args ...interface{}) {
	fmt.Println(append([]interface{}{time.Now().Format(TimeFormat), "[warn]", l.prefix}, args...)...)
}

// Error .
func (l *Stdout) Error(args ...interface{}) {
	fmt.Println(append([]interface{}{time.Now().Format(TimeFormat), "[error]", l.prefix}, args...)...)
}

// Panic .
func (l *Stdout) Panic(args ...interface{}) {
	panic(errors.New(fmt.Sprint(append([]interface{}{time.Now().Format(TimeFormat), "[panic]", l.prefix}, args...)...)))
}

// Fatal .
func (l *Stdout) Fatal(args ...interface{}) {
	fmt.Println(append([]interface{}{time.Now().Format(TimeFormat), "[fatal]", l.prefix}, args...)...)
	os.Exit(-1)
}

// Debugf .
func (l *Stdout) Debugf(template string, args ...interface{}) {
	fmt.Println(time.Now().Format(TimeFormat), "[debug]", l.prefix, fmt.Sprintf(template, args...))
}

// Infof .
func (l *Stdout) Infof(template string, args ...interface{}) {
	fmt.Println(time.Now().Format(TimeFormat), "[info]", l.prefix, fmt.Sprintf(template, args...))
}

// Warnf .
func (l *Stdout) Warnf(template string, args ...interface{}) {
	fmt.Println(time.Now().Format(TimeFormat), "[warn]", l.prefix, fmt.Sprintf(template, args...))
}

// Errorf .
func (l *Stdout) Errorf(template string, args ...interface{}) {
	fmt.Println(time.Now().Format(TimeFormat), "[error]", l.prefix, fmt.Sprintf(template, args...))
}

// Panicf .
func (l *Stdout) Panicf(template string, args ...interface{}) {
	panic(errors.New(fmt.Sprint(append([]interface{}{time.Now().Format(TimeFormat), "[panic]", l.prefix}, args...))))
}

// Fatalf .
func (l *Stdout) Fatalf(template string, args ...interface{}) {
	fmt.Println(time.Now().Format(TimeFormat), "[fatal]", l.prefix, fmt.Sprintf(template, args...))
	os.Exit(-1)
}
