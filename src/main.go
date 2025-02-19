package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/observability-in-deep/lawyers-api/src/config"
	"github.com/observability-in-deep/lawyers-api/src/internal/customer"
	"github.com/observability-in-deep/lawyers-api/src/internal/helthcheck"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initTracer() *trace.TracerProvider {

	config := config.NewConfig()

	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp
}

func main() {
	app := fiber.New()

	tp := initTracer()
	defer func() { _ = tp.Shutdown(context.Background()) }()

	app.Use(otelfiber.Middleware())

	config := config.NewConfig()

	app.Use(logger.New(
		logger.Config{
			Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
			TimeFormat: "15:04:05",
			TimeZone:   "Local",
			Output:     os.Stdout,
		},
	))

	helthcheck.Register(app)
	customer.Register(app)

	err := app.Listen(config.ListenAddress)
	if err != nil {
		log.Fatal(err)
	}

}
