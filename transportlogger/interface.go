package transportlogger

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Printf(string, ...interface{})
	Print(...interface{})
	Println(...interface{})
}

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
