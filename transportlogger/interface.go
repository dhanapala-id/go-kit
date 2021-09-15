package transportlogger

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// Logger is an interface that support log.Logger, logrus.Logger and logrus.Entry
type Logger interface {
	Printf(string, ...interface{})
	Print(...interface{})
	Println(...interface{})
}

// LogrusEntry is an interface that support logrus.Logger and logrus.Entry, this interface
// is used to check whether the passed Logger supports WithField function.
type LogrusEntry interface {
	Logger
	WithField(key string, value interface{}) *logrus.Entry
}

type (
	requestData struct {
		Method string
		URL    string
		Header http.Header
		Body   string
	}
	responseData struct {
		Status int
		Header http.Header
		Body   string
	}
)
