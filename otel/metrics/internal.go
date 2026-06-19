package otelmetrics

import "go.opentelemetry.io/otel/metric"

var ScrapeTime metric.Float64Gauge

func initInternal(meter metric.Meter) (err error) {
	ScrapeTime, err = meter.Float64Gauge("bitcoind_exporter_scrape_time",
		metric.WithDescription("The time it took to scrape the data from bitcoind in milliseconds"),
	)
	return
}
