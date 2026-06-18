package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

type config struct {
	PrometheusEnabled bool `env:"PROMETHEUS_ENABLED" envDefault:"false"`
	PromMetricPort    int `env:"METRIC_PORT" envDefault:"3000"`

	RPCAddress        string `env:"RPC_ADDRESS"`
	RPCUser       string `env:"RPC_USER"`
	RPCPass       string `env:"RPC_PASS"`
	RPCCookieFile string `env:"RPC_COOKIE_FILE"`

	ZmqAddress string `env:"ZMQ_ADDRESS"`

	FetchInterval int `env:"FETCH_INTERVAL" envDefault:"10"`

	OtelEnabled  bool   `env:"OTEL_ENABLED" envDefault:"false"`
	OtelEndpoint string `env:"OTEL_ENDPOINT"`

	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

var (
	log = logrus.WithFields(logrus.Fields{
		"prefix": "config",
	})

	C config
)

func init() {
	loadConfiguration()
}

func loadConfiguration() {
	if config, err := env.ParseAs[config](); err == nil {
		log.Debug("Configuration loaded")
		C = config
	} else {
		log.Fatal(err)
	}

	if C.PrometheusEnabled {
		log.Debug("Prometheus enabled")
	}

	if C.RPCUser == "" && C.RPCPass == "" && C.RPCCookieFile == "" {
		log.Fatal("RPC_USER and RPC_PASS or RPC_COOKIE_FILE must be set")
	}

	if C.OtelEnabled {
		log.Debug("OTel enabled")
		if C.OtelEndpoint == "" {
			log.Fatal("OTEL_ENDPOINT must be set when OTel is enabled")
		}
	}

}
