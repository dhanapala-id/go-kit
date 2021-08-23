package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	im "github.com/dhanapala-id/go-kit/idempotency"
	rs "github.com/dhanapala-id/go-kit/idempotency/store/redis"

	"github.com/go-zoo/bone"
)

func init() {
	im.UseStore(rs.New("127.0.0.1:6379", "", 0))
}

func main() {
	mux := bone.New()

	mux.Get("/check-idempotency", im.Check(http.HandlerFunc(simulateIdempotency)))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func simulateIdempotency(w http.ResponseWriter, r *http.Request) {
	// simulate
	sleep, _ := strconv.Atoi(r.FormValue("sleep"))
	if sleep > 0 {
		time.Sleep(time.Duration(sleep) * time.Second)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
}
