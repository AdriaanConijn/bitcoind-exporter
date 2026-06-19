package otelmetrics

import "go.opentelemetry.io/otel/metric"

var MiningHashrate metric.Float64Gauge

func initMining(meter metric.Meter) (err error) {
	MiningHashrate, err = meter.Float64Gauge("bitcoind_mining_hashrate",
		metric.WithDescription("Mining hashrate in hashes per second"),
	)
	return
}
