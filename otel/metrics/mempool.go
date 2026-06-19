package otelmetrics

import "go.opentelemetry.io/otel/metric"

var (
	MempoolUsage            metric.Float64Gauge
	MempoolMax              metric.Float64Gauge
	MempoolTransactionCount metric.Float64Gauge
)

func initMempool(meter metric.Meter) (err error) {
	if MempoolUsage, err = meter.Float64Gauge("bitcoind_mempool_usage",
		metric.WithDescription("Total memory usage for the mempool in bytes"),
	); err != nil {
		return
	}
	if MempoolMax, err = meter.Float64Gauge("bitcoind_mempool_max",
		metric.WithDescription("Maximum memory usage for the mempool in bytes"),
	); err != nil {
		return
	}
	MempoolTransactionCount, err = meter.Float64Gauge("bitcoind_mempool_transaction_count",
		metric.WithDescription("Total number of transactions in the mempool"),
	)
	return
}
