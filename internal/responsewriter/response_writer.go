package responsewriter

import (
	"bytes"
	"io"
	"net/http"
)

type ResponseWriter struct {
	responseWriter http.ResponseWriter
	statusCode     int
	buffer         *bytes.Buffer
	writer         io.Writer
}

func New(rw http.ResponseWriter) *ResponseWriter {
	b := &bytes.Buffer{}

	return &ResponseWriter{
		responseWriter: rw,
		buffer:         b,
		writer:         io.MultiWriter(([]io.Writer{rw, b})...),
	}
}

func (rw *ResponseWriter) Header() http.Header {
	return rw.responseWriter.Header()
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.responseWriter.WriteHeader(statusCode)
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	return rw.writer.Write(b)
}

func (rw *ResponseWriter) StatusCode() int {
	return rw.statusCode
}

func (rw *ResponseWriter) Body() string {
	return rw.buffer.String()
}
