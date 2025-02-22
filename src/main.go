package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/observability-in-deep/lawyers-api/src/config"
	"github.com/observability-in-deep/lawyers-api/src/internal/customer"
	health "github.com/observability-in-deep/lawyers-api/src/internal/health"
	"github.com/observability-in-deep/lawyers-api/src/internal/lawyers"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func newMeterProvider() (*metric.MeterProvider, error) {

	config := config.NewConfig()
	var metricExporter metric.Exporter
	var err error

	if !config.IsLocal {
		metricExporter, err = otlpmetrichttp.New(
			context.Background(),
			otlpmetrichttp.WithInsecure(),
			otlpmetrichttp.WithEndpoint(config.OtlpEndpoint),
		)
		if err != nil {
			return nil, err
		}
	} else {
		metricExporter, err = stdoutmetric.New()
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			metric.WithInterval(10*time.Second))),
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
		)),
	)
	return meterProvider, nil
}

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
	mp, ok := newMeterProvider()
	if ok != nil {
		log.Fatal(ok)
	}

	defer func() { _ = mp.Shutdown(context.Background()) }()
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

	health.Register(app)
	customer.Register(app)
	lawyers.Register(app)

	err := app.Listen(config.ListenAddress)
	if err != nil {
		log.Println(err)
	}

}
