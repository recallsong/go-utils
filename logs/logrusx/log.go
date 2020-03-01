package logrusx

import (
	"github.com/recallsong/go-utils/logs"
	"github.com/sirupsen/logrus"
)

// Logger .
type Logger struct {
	name string
	*logrus.Entry
}

// New .
func New(options ...interface{}) logs.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	logger := &Logger{"", logrus.NewEntry(log)}
	for _, opt := range options {
		processOptions(logger, opt)
	}
	return logger
}

// Sub .
func (l *Logger) Sub(name string) logs.Logger {
	if len(l.name) > 0 {
		name = l.name + "." + name
	}
	return &Logger{name, l.Entry.WithField("module", name)}
}

func processOptions(logger *Logger, opt interface{}) {
	switch option := opt.(type) {
	case setNameOption:
		logger.name = string(option)
	}
}

type setNameOption string

// WithName .
func WithName(name string) interface{} {
	return setNameOption(name)
}
