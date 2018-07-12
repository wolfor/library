// Log project Log.go
package Log

import (
	mylog "github.com/skoo87/log4go"
)

//外部可访问field
var (
	LoggerCurr *Logger
)

type Logger struct {
}

func NewLogger(logConfigFilePath string) *Logger {
	err := mylog.SetupLogWithConf(logConfigFilePath)

	if err != nil {
		mylog.Fatal("mylog by %s", err)
	}

	return &Logger{}
}

func (l *Logger) Close() {
	mylog.Close()
}

func (l *Logger) Info(fmt string, args ...interface{}) {
	mylog.Info(fmt, args)
}

func (l *Logger) Warn(fmt string, args ...interface{}) {
	mylog.Warn(fmt, args)
}

func (l *Logger) Error(fmt string, args ...interface{}) {
	mylog.Error(fmt, args)
}

func (l *Logger) Debug(fmt string, args ...interface{}) {
	mylog.Debug(fmt, args)
}

func (l *Logger) Fatal(fmt string, args ...interface{}) {
	mylog.Fatal(fmt, args)
}
