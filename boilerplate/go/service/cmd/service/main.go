package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`Hello, Service!`))
	})

	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`{"status": "ok"}`))
	})

	http.ListenAndServe(":9000", nil)
}
