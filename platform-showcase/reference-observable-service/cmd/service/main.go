package main

import (
	"net/http"
	"os"

	"github.com/brpaz/echozap"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"

	"reference-observable-service/internal/logging"
)

func init() {
	_ = godotenv.Load(".platform/env/.env", ".env")
}

func main() {
	e := echo.New()

	logger, loggerSync, _ := logging.NewLogger()
	defer loggerSync()

	e.Use(echozap.ZapLogger(logger))

	c := jaegertracing.New(e, nil)
	defer c.Close()

	e.GET("/", func(c echo.Context) error {
		// zap.L().Info("start processing http request", zap.String("path", c.Path()))
		// defer zap.L().Info("complete processing http request", zap.String("path", c.Path()))

		return c.String(http.StatusOK, "Hello!")
	})

	e.Logger.Fatal(e.Start(os.Getenv("LISTEN_ON")))
}
