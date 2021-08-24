# Go Idempotency Middleware

[![GoDoc](https://pkg.go.dev/badge/github.com/dhanapala-id/go-kit/idempotency)](https://pkg.go.dev/github.com/dhanapala-id/go-kit/idempotency)
![GitHub Action main.yml](https://github.com/dhanapala-id/go-kit/actions/workflows/main.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/dhanapala-id/go-kit/idempotency)](https://goreportcard.com/report/github.com/dhanapala-id/go-kit/idempotency)

A Golang HTTP middleware to make requests idempotent if the client passes an `Idempotency-Key` header.

## How to use

Import the idempotency package and the store you want to use to your code.

```go
import (
	im "github.com/dhanapala-id/go-kit/idempotency"
	rs "github.com/dhanapala-id/go-kit/idempotency/store/redis"
)
```

Set the idempotency middleware to use the store you want.

```go
im.UseStore(rs.New("127.0.0.1:6379", "", 0))
```

Wrap your `http.Handler` using the `Check` function.

```go
mux.Post("/create-order", im.Check(http.HandlerFunc(createOrder)))
```

For a complete example, see the example directory.

## Contributing

We welcome anyone to contribute in this library.
For our contributing guidelines, please check the [CONTRIBUTING.md](https://github.com/dhanapala-id/go-kit/blob/master/CONTRIBUTING.md) file.

## License

This library is under [MIT License](https://choosealicense.com/licenses/mit/), see [LICENSE](https://github.com/dhanapala-id/go-kit/blob/master/LICENSE) file for detail.