package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	service string
)

func init() {
	service = os.Getenv("SERVICE")
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		greeting := fmt.Sprintf("Hello, %s!", service)

		writer.Write([]byte(greeting))
	})

	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`{"status": "ok"}`))
	})

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":9000", nil)
}
