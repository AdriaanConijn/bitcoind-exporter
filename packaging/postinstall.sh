#!/bin/sh
set -e

if ! getent passwd bitcoind-metrics-exporter >/dev/null; then
  useradd --system --no-create-home --shell /usr/sbin/nologin bitcoind-metrics-exporter
fi

systemctl daemon-reload >/dev/null 2>&1 || true
systemctl enable bitcoind-metrics-exporter.service >/dev/null 2>&1 || true

echo "bitcoind-metrics-exporter installed. Edit /etc/bitcoind-metrics-exporter/bitcoind-metrics-exporter.env, then run:"
echo "  systemctl start bitcoind-metrics-exporter"
