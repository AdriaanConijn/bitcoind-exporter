package main

import (
	"runtime"

	"git.aads.cloud/aad/bitcoind-metrics-exporter/config"
	"git.aads.cloud/aad/bitcoind-metrics-exporter/fetcher"
	otelexporter "git.aads.cloud/aad/bitcoind-metrics-exporter/otel"
	"git.aads.cloud/aad/bitcoind-metrics-exporter/zmq"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat:  "2006/01/02 - 15:04:05",
		FullTimestamp:    true,
		QuoteEmptyFields: true,
		SpacePadding:     45,
	})

	log.SetReportCaller(true)

	level, err := log.ParseLevel(config.C.LogLevel)
	if err != nil {
		log.WithError(err).Fatal("Invalid log level")
	}

	log.SetLevel(level)
}

func main() {
	log.WithFields(log.Fields{
		"commit":  commit,
		"runtime": runtime.Version(),
		"arch":    runtime.GOARCH,
	}).Infof("Bitcoind Exporter ₿ %s", version)

	otelexporter.Start()

	go zmq.Start()

	fetcher.Start()
}
