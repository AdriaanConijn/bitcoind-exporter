#!/bin/sh
set -e

systemctl stop bitcoind-exporter.service >/dev/null 2>&1 || true
systemctl disable bitcoind-exporter.service >/dev/null 2>&1 || true
