package otelmetrics

import "go.opentelemetry.io/otel/metric"

var SmartFee metric.Float64Gauge

func initFee(meter metric.Meter) (err error) {
	SmartFee, err = meter.Float64Gauge("bitcoind_smart_fee",
		metric.WithDescription("The estimate fee rate in satoshis per byte"),
	)
	return
}
