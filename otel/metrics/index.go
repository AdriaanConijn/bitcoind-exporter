package otelmetrics

import "go.opentelemetry.io/otel/metric"

var (
	TxIndexSynced     metric.Float64Gauge
	TxIndexBestHeight metric.Float64Gauge
)

func initIndex(meter metric.Meter) (err error) {
	if TxIndexSynced, err = meter.Float64Gauge("bitcoind_txindex_synced",
		metric.WithDescription("Whether the transaction index is synced (1=yes, 0=no)"),
	); err != nil {
		return
	}
	TxIndexBestHeight, err = meter.Float64Gauge("bitcoind_txindex_best_height",
		metric.WithDescription("The best block height in the transaction index"),
	)
	return
}
