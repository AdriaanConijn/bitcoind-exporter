#!/bin/sh
set -e

if ! getent passwd bitcoind-metrics-exporter >/dev/null; then
  useradd --system --no-create-home --shell /usr/sbin/nologin bitcoind-metrics-exporter
fi

systemctl daemon-reload >/dev/null 2>&1 || true
systemctl enable bitcoind-metrics-exporter.service >/dev/null 2>&1 || true

# On upgrade the old prerm left the service running (it only stops on real removal),
# so restart to pick up the new binary. On a fresh install this is a no-op.
if systemctl is-active --quiet bitcoind-metrics-exporter.service; then
  systemctl restart bitcoind-metrics-exporter.service >/dev/null 2>&1 || true
fi

echo "bitcoind-metrics-exporter installed. Edit /etc/bitcoind-metrics-exporter/bitcoind-metrics-exporter.env, then run:"
echo "  systemctl start bitcoind-metrics-exporter"
