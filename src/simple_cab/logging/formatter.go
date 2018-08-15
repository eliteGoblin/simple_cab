package logging

import (
	"github.com/sirupsen/logrus"
	"time"
)

type SimpleCabFormatter struct {
	jsonFormatter *logrus.JSONFormatter
}

func InstanceOfFairFaxFormatter() *SimpleCabFormatter {
	return &SimpleCabFormatter{
		jsonFormatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "_time",
				logrus.FieldKeyLevel: "_level",
			},
			TimestampFormat: time.RFC3339Nano,
		},
	}
}

func (f *SimpleCabFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return f.jsonFormatter.Format(entry)
}
