package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

const logFilePath = "./cache/application.log"

func getTarDayLog() *logrus.Logger {
	log := logrus.New()
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Info("Description Failed to write the log file")
	}
	log.Out = file
	log.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	return log
}

func Info(args ...interface{}) {
	log := getTarDayLog()
	log.Info(args)
}

func InfoF(msg string, args ...interface{}) {
	log := getTarDayLog()
	log.Infof(msg, args)
}

func Error(args ...interface{}) {
	log := getTarDayLog()
	log.Error(args)
}

func ErrorF(msg string, args ...interface{}) {
	log := getTarDayLog()
	log.Info(msg, args)
}

func Debugf(msg string, args ...interface{}) {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logrus.WithFields(logrus.Fields{}).Infof(msg, args)
}

func Debug(args ...interface{}) {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logrus.WithFields(logrus.Fields{}).Info(args)
}
