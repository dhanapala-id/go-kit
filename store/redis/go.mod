module github.com/dhanapala-id/go-idempotency/store/redis

go 1.16

replace github.com/dhanapala-id/go-idempotency/store => ../../store

require (
	github.com/dhanapala-id/go-idempotency/store v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v8 v8.11.3
)
