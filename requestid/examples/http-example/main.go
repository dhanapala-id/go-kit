package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhanapala-id/go-kit/requestid"

	"github.com/go-zoo/bone"
)

func main() {
	mux := bone.New()

	mux.Get("/print-request-id", requestid.Handle(http.HandlerFunc(printRequestID)))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func printRequestID(w http.ResponseWriter, r *http.Request) {
	requestID, referenceID, err := requestid.FromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500 internal server error")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "request id: %s, reference request id: %s", requestID, referenceID)
}
