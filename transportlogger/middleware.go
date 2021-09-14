package transportlogger

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/dhanapala-id/go-kit/responsewriter"
)

type middleware struct {
	log LogrusEntry
}

func NewMiddleare(logger Logger) *middleware {
	llog, ok := logger.(LogrusEntry)
	if !ok {
		llog = log.NewEntry(log.StandardLogger())
	}
	return &middleware{log: llog}
}

func (mw *middleware) Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		tee := io.TeeReader(r.Body, &buf)
		b, _ := ioutil.ReadAll(tee)

		reqData := requestData{
			Method: r.Method,
			URL:    r.URL.String(),
			Header: r.Header,
			Body:   string(b),
		}
		mw.log.WithField("REQUEST_DATA", reqData).Printf("Got request for %s\n", r.URL.String())

		w1, ok := w.(*responsewriter.ResponseWriter)
		if !ok {
			w1 = responsewriter.New(w)
		}
		h.ServeHTTP(w1, r)

		resData := responseData{
			Status: w1.StatusCode(),
			Header: w1.Header().Clone(),
			Body:   w1.Body(),
		}
		mw.log.WithField("RESPONSE_DATA", resData).Printf("Response to client\n")
	})
}
