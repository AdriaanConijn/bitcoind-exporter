package otelexporter

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"git.aads.cloud/aad/bitcoind-metrics-exporter/config"
	otelmetrics "git.aads.cloud/aad/bitcoind-metrics-exporter/otel/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

var log = logrus.WithFields(logrus.Fields{
	"prefix": "otel",
})

func Start() {
	ctx := context.Background()

	var opts []sdkmetric.Option

	if config.C.OtelEndpoint != "" {
		otlpExporter, err := otlpmetrichttp.New(ctx,
			otlpmetrichttp.WithEndpointURL(config.C.OtelEndpoint),
		)
		if err != nil {
			log.WithError(err).Fatal("Failed to create OTLP exporter")
		}

		opts = append(opts, sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(otlpExporter,
				sdkmetric.WithInterval(time.Duration(config.C.FetchInterval)*time.Second),
			),
		))

		log.WithField("endpoint", config.C.OtelEndpoint).Info("OTLP metrics exporter configured.")
	}

	if config.C.PrometheusEnabled {
		promExporter, err := otelprom.New()
		if err != nil {
			log.WithError(err).Fatal("Failed to create Prometheus exporter")
		}

		opts = append(opts, sdkmetric.WithReader(promExporter))
	}

	provider := sdkmetric.NewMeterProvider(opts...)
	otel.SetMeterProvider(provider)

	meter := provider.Meter("bitcoind-metrics-exporter")
	if err := otelmetrics.Init(meter); err != nil {
		log.WithError(err).Fatal("Failed to initialize metrics")
	}

	if config.C.PrometheusEnabled {
		port := strconv.Itoa(config.C.PromMetricPort)
		log.WithField("port", port).Info("Starting Prometheus metrics server.")
		go func() {
			if err := http.ListenAndServe(":"+port, promhttp.Handler()); err != nil {
				log.WithError(err).Error("Prometheus metrics server failed")
			}
		}()
	}
}
