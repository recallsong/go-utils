package logrusx

import (
	"github.com/recallsong/go-utils/logs"
	"github.com/sirupsen/logrus"
)

// // EnableLogContext .
// func EnableLogContext() {
// 	logrus.AddHook(&LogContext{})
// }

// // LogContext .
// type LogContext struct{}

// // Levels .
// func (lc *LogContext) Levels() []logrus.Level {
// 	return logrus.AllLevels
// }

// // Fire .
// func (lc *LogContext) Fire(*logrus.Entry) error {
// 	return nil
// }

// Logger .
type Logger struct {
	name string
	*logrus.Entry
}

// New .
func New(name string) logs.Logger {
	return &Logger{name, logrus.StandardLogger().WithField("module", name)}
}

// Sub .
func (l *Logger) Sub(name string) logs.Logger {
	name = l.name + "." + name
	return &Logger{name, l.Entry.WithField("module", name)}
}
