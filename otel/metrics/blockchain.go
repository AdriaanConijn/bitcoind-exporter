package otelmetrics

import "go.opentelemetry.io/otel/metric"

var (
	BlockchainBlocks               metric.Float64Gauge
	BlockchainHeaders              metric.Float64Gauge
	BlockchainVerificationProgress metric.Float64Gauge
	BlockchainSizeOnDisk           metric.Float64Gauge
)

func initBlockchain(meter metric.Meter) (err error) {
	if BlockchainBlocks, err = meter.Float64Gauge("bitcoind_blockchain_blocks",
		metric.WithDescription("The number of blocks in the blockchain"),
	); err != nil {
		return
	}
	if BlockchainHeaders, err = meter.Float64Gauge("bitcoind_blockchain_headers",
		metric.WithDescription("The number of headers in the blockchain"),
	); err != nil {
		return
	}
	if BlockchainVerificationProgress, err = meter.Float64Gauge("bitcoind_blockchain_verification_progress",
		metric.WithDescription("The verification progress of the blockchain"),
	); err != nil {
		return
	}
	BlockchainSizeOnDisk, err = meter.Float64Gauge("bitcoind_blockchain_size_on_disk",
		metric.WithDescription("The size of the blockchain on disk"),
	)
	return
}
