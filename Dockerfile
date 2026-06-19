FROM gcr.io/distroless/base-nossl-debian13
ARG TARGET_BINARY
COPY ${TARGET_BINARY} /usr/bin/bitcoind-metrics-exporter
ENTRYPOINT ["/usr/bin/bitcoind-metrics-exporter"]
