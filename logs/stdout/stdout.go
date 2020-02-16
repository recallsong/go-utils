package stdout

import (
	"errors"
	"fmt"
	"time"

	"github.com/recallsong/go-utils/log"
)

const TimeFormat = "2006-01-02 15:04:05"

// Stdout .
type Stdout struct {
	name string
}

// Sub .
func (l *Stdout) Sub(name string) log.Logger {
	return &Stdout{
		name: l.name + "." + name,
	}
}

// Debug .
func (l *Stdout) Debug(args ...interface{}) {
	// string len(args)
	// make([]string, )
	[]string{time.Now().Format(TimeFormat), "[debug]", l.name, args...}
	fmt.Println()
}

func (l *Stdout) Info(args ...interface{}) {
	fmt.Println(time.Now().Format(TimeFormat), "[debug]", l.name, args...)
}

func (l *Stdout) Warn(args ...interface{}) {
	fmt.Println(time.Now().Format(TimeFormat), "[debug]", l.name, args...)
}

func (l *Stdout) Error(args ...interface{}) {
	fmt.Println(time.Now().Format(TimeFormat), "[debug]", l.name, args...)
}

func (l *Stdout) Panic(args ...interface{}) {
	errors.New(fmt.Sprint(time.Now().Format(TimeFormat), "[debug]", l.name, args...))
}

func (l *Stdout) Fatal(args ...interface{}) {
	// fmt.Println(time.Now().Format(TimeFormat), "[debug]", l.name, args...)
}

func (l *Stdout) Debugf(template string, args ...interface{}) {

}

func (l *Stdout) Infof(template string, args ...interface{}) {

}

func (l *Stdout) Warnf(template string, args ...interface{}) {

}

func (l *Stdout) Errorf(template string, args ...interface{}) {

}

func (l *Stdout) Panicf(template string, args ...interface{}) {

}

func (l *Stdout) Fatalf(template string, args ...interface{}) {

}
