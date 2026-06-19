#!/bin/sh
set -e

systemctl stop bitcoind-metrics-exporter.service >/dev/null 2>&1 || true
systemctl disable bitcoind-metrics-exporter.service >/dev/null 2>&1 || true
