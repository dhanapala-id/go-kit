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

type Store interface {
	Lock(context.Context, string, time.Duration) error
	Unlock(context.Context, string) error
	Get(context.Context, string) (*Data, error)
	Set(context.Context, string, *Data, time.Duration) error
}

type Data struct {
	Header     http.Header
	StatusCode int
	Body       string
}
