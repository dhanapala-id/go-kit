package main

import (
	"fmt"
	"net/http"

	"github.com/dhanapala-id/go-kit/transportlogger"
	"github.com/go-zoo/bone"
	log "github.com/sirupsen/logrus"
)

func main() {
	mux := bone.New()
	tl := transportlogger.NewMiddleare(log.StandardLogger())

	mux.Post("/log-request", tl.Log(http.HandlerFunc(logRequest)))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func logRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Test", "Hello World!")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world!")
}
