module github.com/dhanapala-id/go-kit/transportlogger/examples/http-example

go 1.16

require (
	github.com/dhanapala-id/go-kit/transportlogger v0.0.0-00010101000000-000000000000
	github.com/go-zoo/bone v1.3.0
	github.com/sirupsen/logrus v1.8.1
)

replace github.com/dhanapala-id/go-kit/transportlogger => ../../

replace github.com/dhanapala-id/go-kit/responsewriter => ../../../responsewriter/
