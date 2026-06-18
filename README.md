# Bitcoind Prometheus Exporter ₿

**Prometheus and OTel metrics for a bitcoin node made simple**

![Buid](https://img.shields.io/github/actions/workflow/status/AdriaanConijn/bitcoind-exporter/release.yml)
![License](https://img.shields.io/github/license/AdriaanConijn/krakendca)

## 🔍 About the project

A Prometheus and OTel Exporter, which provides a deep insight into a Bitcoin full node.

## ⚙️ Configuration

This tool is configured via environment variables. Some environment variables are required and some activate additional functionalities.

| Variable | Description | Required | Default |
| --- | --- | --- | --- |
| `PROMETHEUS_ENABLED` | Whether Prometheus is enabled | ❌ | `true` |
| `RPC_ADDRESS` | The RPC address for the Bitcoin full node, e.g. ``http://127.0.0.1:8332`` | ✅ | |
| `RPC_USER` |The user name that was defined in the Bitcoin Node configuration | ✅ | |
| `RPC_PASS` | The password that was set in the Bitcoin Node configuration | ✅  |  |
| `RPC_COOKIE_FILE` | The path to the cookie file | ✅  |  |
| `ZMQ_ADDRESS` | The address to the ZeroMQ interface of the Bitcoin Fullnode. This variable is required to determine the transcation rates. e.g. ``127.0.0.1:28333`` | ❌ |  |
| `FETCH_INTERVAL` | The interval at which the metrics are to be recalculated. | ❌ | `10` |
| `METRIC_PORT` | The port via which the metrics are provided. | ❌ | `3000` |
| `LOG_LEVEL` | The log level for the service | ❌ | `info` |
| `OTEL_ENABLED` | Whether OpenTelemetry is enabled | ❌ | `false` |
| `OTEL_EXPORTER_OTLP_ENDPOINT` | The OTLP endpoint for the OpenTelemetry collector | ❌ | `otel-collector:4317` |
| `OTEL_SERVICE_NAME` | The service name for the OpenTelemetry collector | ❌ | `bitcoind-exporter` |

Please note that either `RPC_USER` and `RPC_PASS` or `RPC_COOKIE_FILE` must be set.

## 💻 Grafana Dashboard

The official Grafana dashboard can be found here: https://grafana.com/grafana/dashboards/21351

## 🐳 Run with Docker

### Prometheus

#### Docker-CLI

```bash
docker run -d --name bitcoind_exporter \
  -e RPC_ADDRESS=http://127.0.0.1:8332 \
  -e RPC_USER=mempool \
  -e RPC_PASS=mempool \
  -e ZMQ_ADDRESS=127.0.0.1:28333 \
  -e PROMETHEUS_ENABLED=true \
  -v /path/to/cookie/.cookie:/.cookie:ro \
  -p 3000:3000 \
   ghcr.io/AdriaanConijn/bitcoind-exporter:latest
```

#### 🚀 Docker-Compose

```bash
vim docker-compose.yml
```

```yaml
version: "3.8"
services:
  bitcoind_exporter:
    image: ghcr.io/AdriaanConijn/bitcoind-exporter:latest
    ports:
      - "3000:3000"    
    environment:
      - PROMETHEUS_ENABLED=true
      - RPC_ADDRESS=http://127.0.0.1:8332
      - RPC_USER=mempool
      - RPC_PASS=mempool
      - ZMQ_ADDRESS=127.0.0.1:28333
    # Optional, only necessary if Cookie-Auth is to be used
    volumes:
      - /path/to/cookie/.cookie:/.cookie:ro
    restart: always
```

```bash
docker-compose up -d
```

### Otel

#### Docker-CLI

```bash
docker run -d --name bitcoind_exporter \
  -e RPC_ADDRESS=http://127.0.0.1:8332 \
  -e RPC_USER=mempool \
  -e RPC_PASS=mempool \
  -e ZMQ_ADDRESS=127.0.0.1:28333 \
  -e OTEL_ENABLED=true \
  -e OTEL_ENDPOINT= http://localhost:4123 \
  -v /path/to/cookie/.cookie:/.cookie:ro \
   ghcr.io/AdriaanConijn/bitcoind-exporter:latest
```

#### 🚀 Docker-Compose

```bash
vim docker-compose.yml
```

```yaml
version: "3.8"
services:
  bitcoind_exporter:
    image: ghcr.io/AdriaanConijn/bitcoind-exporter:latest
    environment:
      - OTEL_ENABLED=true
      - OTEL_ENDPOINT= http://localhost:4123
      - RPC_ADDRESS=http://127.0.0.1:8332
      - RPC_USER=mempool
      - RPC_PASS=mempool
      - ZMQ_ADDRESS=127.0.0.1:28333
    # Optional, only necessary if Cookie-Auth is to be used
    volumes:
      - /path/to/cookie/.cookie:/.cookie:ro
    restart: always
```

```bash
docker-compose up -d
```
