package log

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger = NewLogrusLogger()

func SetUpLogPath(logpath string) {
	if logpath != "stdout" {
		file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("unable init logging: %v", err)
		}
		Log.SetOutput(file)
		defer file.Close()
	}
}

func Error(msg string, cause string) {
	Log.WithFields(logrus.Fields{
		"cause": cause,
	}).Error(msg)
}

func Info(msg string) {
	Log.Info(msg)
}

func Debug(msg string) {
	Log.Debug(msg)
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
