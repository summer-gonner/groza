package middleware

import (
	"github.com/sirupsen/logrus"
	"time"
)

var Logging *logrus.Entry

func InitLogging() {
	baseLogger := logrus.New()
	baseLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.DateTime,
		FullTimestamp:   true,
	})
	Logging = logrus.NewEntry(baseLogger)

}
