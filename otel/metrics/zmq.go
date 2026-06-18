package otelmetrics

import "go.opentelemetry.io/otel/metric"

var ZmqTransactions metric.Int64Counter

func initZmq(meter metric.Meter) (err error) {
	ZmqTransactions, err = meter.Int64Counter("bitcoind_transactions_total",
		metric.WithDescription("Total number of transactions observed via ZMQ"),
	)
	return
}
