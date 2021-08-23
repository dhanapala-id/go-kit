module github.com/dhanapala-id/go-idempotency/examples/http-example

go 1.16

require (
	github.com/dhanapala-id/go-idempotency v0.0.0-00010101000000-000000000000
	github.com/dhanapala-id/go-idempotency/store/redis v0.0.0-00010101000000-000000000000
	github.com/go-zoo/bone v1.3.0
)

replace github.com/dhanapala-id/go-idempotency => ../../

replace github.com/dhanapala-id/go-idempotency/store => ../../store

replace github.com/dhanapala-id/go-idempotency/store/redis => ../../store/redis
