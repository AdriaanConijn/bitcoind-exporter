package otelmetrics

import "go.opentelemetry.io/otel/metric"

func Init(meter metric.Meter) error {
	for _, fn := range []func(metric.Meter) error{
		initBlockchain,
		initMempool,
		initMemory,
		initIndex,
		initNetwork,
		initFee,
		initMining,
		initInternal,
		initZmq,
	} {
		if err := fn(meter); err != nil {
			return err
		}
	}
	return nil
}
