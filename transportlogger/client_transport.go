package transportlogger

import (
	"bytes"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type roundTripper struct {
	transport http.RoundTripper
	log       LogrusEntry
}

// NewRoundTripper creates a new round tripper instance that implements http.RoundTripper
// with input log.Logger, logrus.Logger or logrus.Entry.
func NewRoundTripper(rt http.RoundTripper, logger Logger) *roundTripper {
	llog, ok := logger.(LogrusEntry)
	if !ok {
		llog = log.NewEntry(log.StandardLogger())
	}
	return &roundTripper{rt, llog}
}

// RoundTrip intercepts the request and log the request and response to the logger.
func (rt roundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	var reqBytes []byte
	if req.Body != nil {
		reqBytes, _ = ioutil.ReadAll(req.Body)
		err = req.Body.Close()
		if err != nil {
			rt.log.Printf("Error: %v", err)
			return
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBytes))
	}

	reqData := requestData{
		Method: req.Method,
		URL:    req.URL.String(),
		Header: req.Header.Clone(),
		Body:   string(reqBytes),
	}
	rt.log.WithField("REQUEST_DATA", reqData).Printf("Sending request to %v\n", req.URL)

	res, err = rt.transport.RoundTrip(req)
	if err != nil {
		rt.log.Printf("Error: %v", err)
		return
	}

	var resBytes []byte
	if res.Body != nil {
		resBytes, _ = ioutil.ReadAll(res.Body)
		err = res.Body.Close()
		if err != nil {
			rt.log.Printf("Error: %v", err)
			return
		}
		res.Body = ioutil.NopCloser(bytes.NewBuffer(resBytes))
	}

	resData := responseData{
		Status: res.StatusCode,
		Header: res.Header.Clone(),
		Body:   string(resBytes),
	}
	rt.log.WithField("RESPONSE_DATA", resData).Println("Got response")

	return
}
