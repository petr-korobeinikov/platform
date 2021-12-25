package main

import (
	"os"
	"strconv"

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

	// Provide a library to handle env vars: env.E("FOO").AsUint(); env.E("FOO").AsDuration()
	randomGeneratorMinValue, _ := strconv.Atoi(os.Getenv("RANDOM_GENERATOR_MIN_NUMBER"))
	randomGeneratorMaxValue, _ := strconv.Atoi(os.Getenv("RANDOM_GENERATOR_MAX_NUMBER"))

	fibonacciCountingServiceMaxFibNumber, _ := strconv.Atoi(os.Getenv("FIBONACCI_COUNTING_SERVICE_MAX_N_NUMBER"))

	randomGenerator := random.NewGenerator(randomGeneratorMinValue, randomGeneratorMaxValue)
	fibonacciCountingService := fibonacci.NewCountingService(fibonacciCountingServiceMaxFibNumber)

	indexHandler := handler.NewIndexHandler()
	complexHandler := handler.NewComplexHandler(handler.ComplexHandlerCfg{
		RandomGenerator:          randomGenerator,
		FibonacciCountingService: fibonacciCountingService,
	})

	e.GET("/", indexHandler.HandleRequest)
	e.GET("/complex", complexHandler.HandleRequest)

	e.Logger.Fatal(e.Start(os.Getenv("LISTEN_ON")))
}
