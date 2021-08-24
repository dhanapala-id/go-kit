package responsewriter

import (
	"bytes"
	"io"
	"net/http"
)

// ResponseWriter implements the http.ResponseWriter,
// supports for writting the response to a buffer for
// later can be used to cache the response.
type ResponseWriter struct {
	responseWriter http.ResponseWriter
	statusCode     int
	buffer         *bytes.Buffer
	writer         io.Writer
}

// New creates a new ResponseWriter based on the handler ResponseWriter
func New(rw http.ResponseWriter) *ResponseWriter {
	b := &bytes.Buffer{}

	return &ResponseWriter{
		responseWriter: rw,
		buffer:         b,
		writer:         io.MultiWriter(([]io.Writer{rw, b})...),
	}
}

// Header returns the current set headers.
func (rw *ResponseWriter) Header() http.Header {
	return rw.responseWriter.Header()
}

// WriteHeader writes the header to the response and stores the status code for later use.
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.responseWriter.WriteHeader(statusCode)
}

// Write writes the body to the handler ResponseWriter and the buffer.
func (rw *ResponseWriter) Write(b []byte) (int, error) {
	return rw.writer.Write(b)
}

// StatusCode returns the status code that was written to the ResponseWriter.
func (rw *ResponseWriter) StatusCode() int {
	return rw.statusCode
}

// Body returns the response body that was written to the buffer.
func (rw *ResponseWriter) Body() string {
	return rw.buffer.String()
}
