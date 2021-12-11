package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var (
	service string
)

func init() {
	_ = godotenv.Load(".platform/env/.env", ".env")

	service = os.Getenv("SERVICE")
}

func main() {
	logCfg := zap.NewProductionConfig()
	logCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := logCfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		zap.L().Info("go http request", zap.String("address", "/"))

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
