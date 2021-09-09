package requestid

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type contextKey struct{}

var ctxKey = &contextKey{}

const (
	RequestIDHeader          = "X-Request-ID"
	ReferenceRequestIDHeader = "X-Reference-Request-ID"
)

var ErrNoRequestID = errors.New("requestid: no requestID found in context")

type reqID struct {
	RequestID   string
	ReferenceID string
}

func generateID(requestID, referenceID string) (string, string) {
	if requestID == "" {
		requestID = uuid.New().String()
	}

	if referenceID == "" {
		referenceID = requestID
	}

	return requestID, referenceID
}

func setContext(ctx context.Context, requestID, referenceID string) context.Context {
	return context.WithValue(ctx, ctxKey, reqID{requestID, referenceID})
}

// FromContext returns the request id and reference request id value.
// The input is a context that has been set the request id and reference request id.
// It will return empty values and an error if no request id was set.
func FromContext(ctx context.Context) (string, string, error) {
	val := ctx.Value(ctxKey)
	if val == nil {
		return "", "", ErrNoRequestID
	}
	id, _ := val.(reqID)
	return id.RequestID, id.ReferenceID, nil
}

// Handle is the HTTP middleware to set a request id to the response context.
// It sets the request id and reference request id from X-Request-ID and X-Reference-Request-ID header respectively.
// To get the current request id and reference request id, use FromContext and pass the http.Request context.
func Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID, referenceID := generateID(
			r.Header.Get(RequestIDHeader),
			r.Header.Get(ReferenceRequestIDHeader),
		)

		ctx := setContext(r.Context(), requestID, referenceID)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
