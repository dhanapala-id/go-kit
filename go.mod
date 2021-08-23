module github.com/dhanapala-id/go-idempotency

go 1.16

require github.com/go-redis/redis/v8 v8.11.3

replace github.com/dhanapala-id/go-idempotency/internal/store/redis => ./internal/store/redis

replace github.com/dhanapala-id/go-idempotency/internal/store => ./internal/store
