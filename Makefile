.PHONY: build
build:
	CGO_ENABLED=0 go build -o prometheus-grafana-demo cmd/main.go

