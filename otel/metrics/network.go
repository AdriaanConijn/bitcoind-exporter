package otelmetrics

import "go.opentelemetry.io/otel/metric"

var (
	TotalConnections metric.Float64Gauge
	ConnectionsIn    metric.Float64Gauge
	ConnectionsOut   metric.Float64Gauge
	TotalBytesRecv   metric.Float64Gauge
	TotalBytesSent   metric.Float64Gauge
)

func initNetwork(meter metric.Meter) (err error) {
	if TotalConnections, err = meter.Float64Gauge("bitcoind_connections_total",
		metric.WithDescription("Total number of peer connections"),
	); err != nil {
		return
	}
	if ConnectionsIn, err = meter.Float64Gauge("bitcoind_connections_in",
		metric.WithDescription("Number of inbound peer connections"),
	); err != nil {
		return
	}
	if ConnectionsOut, err = meter.Float64Gauge("bitcoind_connections_out",
		metric.WithDescription("Number of outbound peer connections"),
	); err != nil {
		return
	}
	if TotalBytesRecv, err = meter.Float64Gauge("bitcoind_total_bytes_recv",
		metric.WithDescription("Total bytes received from peers"),
	); err != nil {
		return
	}
	TotalBytesSent, err = meter.Float64Gauge("bitcoind_total_bytes_sent",
		metric.WithDescription("Total bytes sent to peers"),
	)
	return
}
