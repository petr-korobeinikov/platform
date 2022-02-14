package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/pkorobeinikov/environ"

	"service-docker-compose-generator/internal/cfg/dotenv"
)

func init() {
	dotenv.Load(".env", ".platform/env/.env")
}

func main() {
	service, err := environ.E("SERVICE").AsString()
	if err != nil {
		println("env SERVICE is not defined")
		return
	}

	httpApiPort, err := environ.E("HTTP_API_PORT").AsInt()
	if err != nil {
		println("env HTTP_API_PORT is not defined")
		return
	}

	e := echo.New()

	p := prometheus.NewPrometheus(strings.ReplaceAll(service, "-", "_"), nil)
	p.Use(e)

	c := jaegertracing.New(e, nil)
	defer c.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Hello, I am %s!", service))
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", httpApiPort)))
}
