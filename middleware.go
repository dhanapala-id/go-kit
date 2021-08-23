package idempotency

import (
	"net/http"
	"time"

	"github.com/dhanapala-id/go-idempotency/internal/responsewriter"
	"github.com/dhanapala-id/go-idempotency/internal/store"
)

const (
	defaultKeyName         string        = "Idempotency-Key"
	defaultKeyExpiration   time.Duration = 5 * time.Minute
	defaultCacheExpiration time.Duration = 1 * time.Hour
)

var defaultInstance *Idempotency = New()

type (
	Idempotency struct {
		store           store.Store
		keyName         string
		keyExpiration   time.Duration
		cacheExpiration time.Duration
	}
	Option struct {
		KeyName         *string
		KeyExpiration   *time.Duration
		CacheExpiration *time.Duration
	}
)

// New initiates a new idempotency middleware instance using default configuration.
func New() *Idempotency {
	return &Idempotency{
		keyName:         defaultKeyName,
		keyExpiration:   defaultKeyExpiration,
		cacheExpiration: defaultCacheExpiration,
	}
}

// NewWithOption initiates a new idempotency middleware instance.
func NewWithOption(opt Option) *Idempotency {
	o := New()
	o.UseOption(opt)
	return o
}

// UseOption sets the middleware instance configuration.
func (o *Idempotency) UseOption(opt Option) {
	if opt.KeyName != nil {
		o.keyName = *opt.KeyName
	}
	if opt.KeyExpiration != nil {
		o.keyExpiration = *opt.KeyExpiration
	}
	if opt.CacheExpiration != nil {
		o.cacheExpiration = *opt.CacheExpiration
	}
}

// UseStore sets the middleware instance store.
func (o *Idempotency) UseStore(s store.Store) {
	o.store = s
}

// Check is the instance middleware, it checks if the HTTP header has an idempotency key header
// and return the cached response when the same previous request already done processing.
// If the check failes to retrieve lock from the store, it will return `409 Conflict`.
func (o *Idempotency) Check(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get(o.keyName)
		if key == "" {
			h.ServeHTTP(w, r)
			return
		}

		// check if store is initialized
		if o.store == nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		ctx := r.Context()

		// get cached data
		data, err := o.store.Get(ctx, key)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// if cached data not empty, return from cache
		if data != nil {
			for k, v := range data.Header {
				for k1 := range v {
					w.Header().Add(k, v[k1])
				}
			}
			w.WriteHeader(data.StatusCode)
			w.Write([]byte(data.Body))
			return
		}

		// lock key
		if err := o.store.Lock(ctx, key, o.keyExpiration); err != nil {
			http.Error(w, "conflict", http.StatusConflict)
			return
		}

		// create new response writer to support cache response body
		w1 := responsewriter.New(w)
		h.ServeHTTP(w1, r)

		// cache data
		o.store.Set(ctx, key, &store.Data{
			StatusCode: w1.StatusCode(),
			Header:     w1.Header().Clone(),
			Body:       w1.Body(),
		}, o.cacheExpiration)

		// unlock key
		o.store.Unlock(ctx, key)
	})
}

// UseOption sets the default middleware instance configuration.
func UseOption(opt Option) {
	defaultInstance.UseOption(opt)
}

// UseStore sets the default middleware instance store.
func UseStore(s store.Store) {
	defaultInstance.UseStore(s)
}

// Check is the default instance middleware.
func Check(h http.Handler) http.Handler {
	return defaultInstance.Check(h)
}
