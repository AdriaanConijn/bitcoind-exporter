package otelmetrics

import "go.opentelemetry.io/otel/metric"

var (
	TotalConnections    metric.Float64Gauge
	ConnectionsIn       metric.Float64Gauge
	ConnectionsOut      metric.Float64Gauge
	TotalBytesRecv      metric.Float64Gauge
	TotalBytesSent      metric.Float64Gauge
	Ipv4ConnectionsIn   metric.Float64Gauge
	Ipv6ConnectionsIn   metric.Float64Gauge
	OnionConnectionsIn  metric.Float64Gauge
	Ipv4ConnectionsOut  metric.Float64Gauge
	Ipv6ConnectionsOut  metric.Float64Gauge
	OnionConnectionsOut metric.Float64Gauge
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
	if TotalBytesSent, err = meter.Float64Gauge("bitcoind_total_bytes_sent",
		metric.WithDescription("Total bytes sent to peers"),
	); err != nil {
		return
	}
	if Ipv4ConnectionsIn, err = meter.Float64Gauge("bitcoind_ipv4_connections_in",
		metric.WithDescription("Number of IPv4 connections"),
	); err != nil {
		return
	}
	if Ipv6ConnectionsIn, err = meter.Float64Gauge("bitcoind_ipv6_connections_in",
		metric.WithDescription("Number of IPv6 connections"),
	); err != nil {
		return
	}
	if OnionConnectionsIn, err = meter.Float64Gauge("bitcoind_onion_connections_in",
		metric.WithDescription("Number of Onion connections"),
	); err != nil {
		return
	}
	if Ipv4ConnectionsOut, err = meter.Float64Gauge("bitcoind_ipv4_connections_out",
		metric.WithDescription("Number of IPv4 connections"),
	); err != nil {
		return
	}
	if Ipv6ConnectionsOut, err = meter.Float64Gauge("bitcoind_ipv6_connections_out",
		metric.WithDescription("Number of IPv6 connections"),
	); err != nil {
		return
	}
	if OnionConnectionsOut, err = meter.Float64Gauge("bitcoind_onion_connections_out",
		metric.WithDescription("Number of Onion connections"),
	); err != nil {
		return
	}

	return
}
