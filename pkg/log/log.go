package log

import (
	"log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once
var logger *logrus.Logger

func Instance() *logrus.Logger {
	once.Do(func() {
		if logger == nil {
			logger = NewLogrusLogger()
		}
	})
	return logger
}

func SetUpLogPath(logpath string) *os.File {
	if logpath != "stdout" {
		file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("unable init logging: %v", err)
		}
		Instance().SetOutput(file)
		return file
	}
	return nil
}

func Error(msg string, cause string) {
	Instance().WithFields(logrus.Fields{
		"cause": cause,
	}).Error(msg)
}

func Info(msg string) {
	Instance().Info(msg)
}

func Debug(msg string) {
	Instance().Debug(msg)
}

func Fatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}

func NewLogrusLogger() *logrus.Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})
	logrusLogger.SetLevel(logrus.DebugLevel)
	return logrusLogger
}
