module github.com/dhanapala-id/go-kit/idempotency/examples/http-example

go 1.16

require (
	github.com/dhanapala-id/go-kit/idempotency v0.0.0-00010101000000-000000000000
	github.com/dhanapala-id/go-kit/idempotency/store/redis v0.0.0-00010101000000-000000000000
	github.com/go-zoo/bone v1.3.0
)

replace github.com/dhanapala-id/go-kit/idempotency => ../../

replace github.com/dhanapala-id/go-kit/idempotency/store/redis => ../../store/redis
