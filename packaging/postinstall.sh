#!/bin/sh
set -e

if ! getent passwd bitcoind-exporter >/dev/null; then
  useradd --system --no-create-home --shell /usr/sbin/nologin bitcoind-exporter
fi

systemctl daemon-reload >/dev/null 2>&1 || true
systemctl enable bitcoind-exporter.service >/dev/null 2>&1 || true

echo "bitcoind-exporter installed. Edit /etc/bitcoind-exporter/bitcoind-exporter.env, then run:"
echo "  systemctl start bitcoind-exporter"
