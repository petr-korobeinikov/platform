package main

import (
	"os"

	"github.com/brpaz/echozap"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"

	"reference-observable-service/internal/handler"
	"reference-observable-service/internal/observability/logging"
	"reference-observable-service/internal/service/fibonacci"
	"reference-observable-service/internal/service/random"
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

	randomGenerator := random.NewGenerator(-3, 10)
	fibonacciCountingService := fibonacci.NewCountingService(7)

	indexHandler := handler.NewIndexHandler()
	complexHandler := handler.NewComplexHandler(handler.ComplexHandlerCfg{
		RandomGenerator:          randomGenerator,
		FibonacciCountingService: fibonacciCountingService,
	})

	e.GET("/", indexHandler.HandleRequest)
	e.GET("/complex", complexHandler.HandleRequest)

	e.Logger.Fatal(e.Start(os.Getenv("LISTEN_ON")))
}
