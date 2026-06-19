package otelmetrics

import "go.opentelemetry.io/otel/metric"

var (
	MemoryUsed metric.Float64Gauge
	MemoryFree metric.Float64Gauge
	MemoryTotal metric.Float64Gauge
	MemoryLocked metric.Float64Gauge
	ChunksUsed metric.Float64Gauge
	ChunksFree metric.Float64Gauge
)

func initMemory(meter metric.Meter) (err error) {
	if MemoryUsed, err = meter.Float64Gauge("bitcoind_memory_used",
		metric.WithDescription("Used locked memory in bytes"),
	); err != nil {
		return
	}
	if MemoryFree, err = meter.Float64Gauge("bitcoind_memory_free",
		metric.WithDescription("Free locked memory in bytes"),
	); err != nil {
		return
	}
	if MemoryTotal, err = meter.Float64Gauge("bitcoind_memory_total",
		metric.WithDescription("Total locked memory in bytes"),
	); err != nil {
		return
	}
	if MemoryLocked, err = meter.Float64Gauge("bitcoind_memory_locked",
		metric.WithDescription("Amount of memory that is locked"),
	); err != nil {
		return
	}
	if ChunksUsed, err = meter.Float64Gauge("bitcoind_memory_chunks_used",
		metric.WithDescription("Number of allocated memory chunks"),
	); err != nil {
		return
	}
	ChunksFree, err = meter.Float64Gauge("bitcoind_memory_chunks_free",
		metric.WithDescription("Number of unused memory chunks"),
	)
	return
}
