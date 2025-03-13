GOCMD=GO111MODULE=on CGO_ENABLED=0 go
GOBUILD=${GOCMD} build

.PHONY: init
# Initialize environment
init:
	go install github.com/google/wire/cmd/wire@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest

.PHONY: generate
# Generate all
generate: proto api

.PHONY: version
# Show the generated version
version:
	@find app -type d -depth 1 -print | xargs -L 1 bash -c 'echo "version: $$0" && cd "$$0" && $(MAKE) version'

.PHONY: wire
# Generate wire
wire:
	@find app -type d -depth 1 -print | xargs -L 1 bash -c 'echo "wire: $$0" && cd "$$0" && $(MAKE) wire'

.PHONY: proto
# Generate internal proto struct
proto:
	buf generate

.PHONY: api
# Generate API
api:
	buf generate --template=buf.gen.client.yaml
	buf generate --template=buf.gen.server.yaml

.PHONY: build
# Build executable file
build:
	@find app -type d -depth 1 -print | xargs -L 1 bash -c 'echo "build: $$0" && cd "$$0" && $(MAKE) build'

.PHONY: run
# Start all project services
run: start log

.PHONY: start
# Start all project services
start: stop build
	@find app -type d -depth 1 -print | xargs -L 1 bash -c 'echo "start: $$0" && cd "$$0" && $(MAKE) start'

.PHONY: stop
# Stop running project services
stop:
	@find app -type d -depth 1 -print | xargs -L 1 bash -c 'echo "stop: $$0" && cd "$$0" && $(MAKE) stop'

.PHONY: log
# tail -f app/gate/bin/debug.log
log:
	@find app -type d -depth 1 -print | xargs -L 1 bash -c 'echo "log: $$0" && cd "$$0" && tail -f bin/debug.log'

.PHONY: test
# Run tests
test:
	go test -v ./... -cover
# Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
