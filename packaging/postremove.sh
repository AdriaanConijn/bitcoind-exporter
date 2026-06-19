#!/bin/sh
set -e

systemctl daemon-reload >/dev/null 2>&1 || true

if [ "$1" = "purge" ]; then
  userdel bitcoind-exporter >/dev/null 2>&1 || true
fi
