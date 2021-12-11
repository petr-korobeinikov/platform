package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	service string
)

func init() {
	service = os.Getenv("SERVICE")

	_ = godotenv.Load(".platform/env/.env", ".env")
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-type", "text/plain")

		greeting := fmt.Sprintf("Hello, %s!\n", service)

		writer.Write([]byte(greeting))

		writer.Write([]byte("\n\nMy environment is:\n"))
		for _, s := range os.Environ() {
			writer.Write([]byte(s + "\n"))
		}
	})

	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`{"status": "ok"}`))
	})

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":9000", nil)
}
