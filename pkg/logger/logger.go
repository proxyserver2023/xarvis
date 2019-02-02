package logger

import (
	"github.com/sirupsen/logrus"
)

// Logger defines the logger interface that is exposed via RequestScope.
type Logger interface {
	// adds a field that should be added to every message being logged
	SetField(name, value string)

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

type mylogger struct {
	logger *logrus.Logger
	fields logrus.Fields
}

func NewLogger(l *logrus.Logger, fields logrus.Fields) Logger {
	return &mylogger{
		logger: l,
		fields: fields,
	}
}

func (l *mylogger) SetField(name, value string) {
	l.fields[name] = value
}

func (l *mylogger) Debugf(format string, args ...interface{}) {
	l.tagged().Debugf(format, args...)
}

func (l *mylogger) Infof(format string, args ...interface{}) {
	l.tagged().Infof(format, args...)
}

func (l *mylogger) Warnf(format string, args ...interface{}) {
	l.tagged().Warnf(format, args...)
}

func (l *mylogger) Errorf(format string, args ...interface{}) {
	l.tagged().Errorf(format, args...)
}

func (l *mylogger) Debug(args ...interface{}) {
	l.tagged().Debug(args...)
}

func (l *mylogger) Info(args ...interface{}) {
	l.tagged().Info(args...)
}

func (l *mylogger) Warn(args ...interface{}) {
	l.tagged().Warn(args...)
}

func (l *mylogger) Error(args ...interface{}) {
	l.tagged().Error(args...)
}

func (l *mylogger) tagged() *logrus.Entry {
	return l.logger.WithFields(l.fields)
}
