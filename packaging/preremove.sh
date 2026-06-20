#!/bin/sh
set -e

# $1 is "upgrade" during apt upgrade and "remove"/"deconfigure" on actual removal.
# Only stop/disable on real removal so upgrades leave the service running.
if [ "$1" = "remove" ] || [ "$1" = "deconfigure" ]; then
  systemctl stop bitcoind-metrics-exporter.service >/dev/null 2>&1 || true
  systemctl disable bitcoind-metrics-exporter.service >/dev/null 2>&1 || true
fi
