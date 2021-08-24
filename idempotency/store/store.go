package store

import (
	"context"
	"errors"
	"net/http"
	"time"
)

var (
	ErrUnableToLock = errors.New("unable to acquire lock")
)

// Store is the interface to implement a different kind of storage strategy.
type Store interface {
	Lock(context.Context, string, time.Duration) error
	Unlock(context.Context, string) error
	Get(context.Context, string) (*Data, error)
	Set(context.Context, string, *Data, time.Duration) error
}

// Data is used to fetch and store data in the storage.
// Contains the response status code, headers and body.
type Data struct {
	Header     http.Header
	StatusCode int
	Body       string
}
