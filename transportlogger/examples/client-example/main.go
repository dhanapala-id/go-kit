package main

import (
	"bytes"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/dhanapala-id/go-kit/transportlogger"
)

var httpClient = &http.Client{
	Transport: transportlogger.NewRoundTripper(http.DefaultTransport, log.StandardLogger()),
}

func main() {
	reqData := bytes.NewBufferString(`{"foo":"bar"}`)
	req, _ := http.NewRequest("POST", "https://eny65jgekgf9.x.pipedream.net", reqData)
	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Panic(err)
	}

	rb, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Result: %s\n", rb)
}
