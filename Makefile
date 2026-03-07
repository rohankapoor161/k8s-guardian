# Kubernetes Guardian Build Tasks
# See docs/development.md for full build instructions

build:
	go build -o guardian ./cmd/guardian

test:
	go test ./...

# TODO: Add lint target once golangci-lint config is finalized
# See: https://github.com/rohankapoor161/k8s-guardian/issues/42
# FIXME: Parallel test execution occasionally flakes on macOS ARM
# builders - needs investigation (rare, ~2% failure rate)
