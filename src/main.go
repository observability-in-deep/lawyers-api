package main

import (
	"context"
	"log"
	"os"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/observability-in-deep/lawyers-api/src/config"
	"github.com/observability-in-deep/lawyers-api/src/internal/customer"
	health "github.com/observability-in-deep/lawyers-api/src/internal/health"
	"github.com/observability-in-deep/lawyers-api/src/internal/lawyers"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initTracer() *trace.TracerProvider {

	config := config.NewConfig()
	var exporter trace.SpanExporter
	var err error

	if !config.IsLocal {
		exporter, err = otlptracehttp.New(
			context.Background(),
			otlptracehttp.WithInsecure(),
			otlptracehttp.WithEndpoint(config.OtlpEndpoint),
		)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		exporter, err = stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			log.Fatal(err)
		}
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

	prometheus := fiberprometheus.New("lawyers-api")

	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/health" {
			return c.Next()
		}
		return otelfiber.Middleware()(c)
	})

	config := config.NewConfig()

	app.Use(logger.New(
		logger.Config{
			Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error} | " + config.ServiceName + "\n",
			TimeFormat: "15:04:05",
			TimeZone:   "Local",
			Output:     os.Stdout,
		},
	))

	health.Register(app)
	customer.Register(app)
	lawyers.Register(app)

	err := app.Listen(config.ListenAddress)
	if err != nil {
		log.Println(err)
	}

}
