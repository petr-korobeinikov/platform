package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`Hello, Service!`))
	})

	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`{"status": "ok"}`))
	})

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":9000", nil)
}
