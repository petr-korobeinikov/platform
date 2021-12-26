package main

import (
	"strings"

	"github.com/brpaz/echozap"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/pkorobeinikov/environ"
	"go.uber.org/zap"

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

	serviceName, err := environ.E("SERVICE").AsString()
	if err != nil {
		logger.Fatal("service name not provided", zap.Error(err))
	}

	p := prometheus.NewPrometheus(strings.ReplaceAll(serviceName, "-", "_"), nil)
	p.Use(e)

	c := jaegertracing.New(e, nil)
	defer c.Close()

	randomGeneratorMinValue, _ := environ.E("RANDOM_GENERATOR_MIN_NUMBER").Default(-3).AsInt()
	randomGeneratorMaxValue, _ := environ.E("RANDOM_GENERATOR_MAX_NUMBER").Default(10).AsInt()

	fibonacciCountingServiceMaxFibNumber, _ := environ.E("FIBONACCI_COUNTING_SERVICE_MAX_N_NUMBER").
		Default(7).
		AsInt()

	randomGenerator := random.NewGenerator(randomGeneratorMinValue, randomGeneratorMaxValue)
	fibonacciCountingService := fibonacci.NewCountingService(fibonacciCountingServiceMaxFibNumber)

	indexHandler := handler.NewIndexHandler()
	complexHandler := handler.NewComplexHandler(handler.ComplexHandlerCfg{
		RandomGenerator:          randomGenerator,
		FibonacciCountingService: fibonacciCountingService,
	})

	e.GET("/", indexHandler.HandleRequest)
	e.GET("/complex", complexHandler.HandleRequest)

	listenOn, _ := environ.E("LISTEN_ON").Default(":9000").AsString()
	e.Logger.Fatal(e.Start(listenOn))
}
