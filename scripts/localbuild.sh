go run github.com/goreleaser/goreleaser/v2@latest release --snapshot --clean --skip=publish

IMG=localhost/bitcoind-metrics-exporter
VERSION=0.0.0-test
AMD64_BIN=$(find dist -type f -name bitcoind-metrics-exporter -path '*linux_amd64*' | head -n1)
ARM64_BIN=$(find dist -type f -name bitcoind-metrics-exporter -path '*linux_arm64*' | head -n1)
echo "amd64: $AMD64_BIN"
echo "arm64: $ARM64_BIN"
cp "$AMD64_BIN" ./bitcoind-metrics-exporter.amd64
cp "$ARM64_BIN" ./bitcoind-metrics-exporter.arm64
buildah build --arch amd64 --build-arg TARGET_BINARY=bitcoind-metrics-exporter.amd64 -t "$IMG:$VERSION-amd64" .
