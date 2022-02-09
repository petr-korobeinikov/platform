package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func init() {
	_ = godotenv.Load(".platform/env/.env", ".env")
}

func main() {
	e := echo.New()

	c := jaegertracing.New(e, nil)
	defer c.Close()

	p := prometheus.NewPrometheus(strings.ReplaceAll(os.Getenv("SERVICE"), "-", "_"), nil)
	p.Use(e)

	e.GET("/", func(c echo.Context) error {
		// Use templating on the next step.
		// https://echo.labstack.com/guide/templates/
		var sb strings.Builder

		sb.WriteString("Service Desktop\n\n")

		if _, enabled := os.LookupEnv("COMPONENT_JAEGERUI_ENABLED"); enabled {
			sb.WriteString("* Jaeger UI: ")
			sb.WriteString(os.Getenv("COMPONENT_JAEGERUI_HOST"))
			sb.WriteString("\n")
		}

		if _, enabled := os.LookupEnv("COMPONENT_KAFDROP_ENABLED"); enabled {
			sb.WriteString("* Kafdrop: ")
			sb.WriteString(os.Getenv("COMPONENT_KAFDROP_HOST"))
			sb.WriteString("\n")
		}

		return c.String(http.StatusOK, sb.String())
	})

	e.Logger.Fatal(e.Start(os.Getenv("LISTEN_ON")))
}
