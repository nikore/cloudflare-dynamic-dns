FILES = $$(find . -type f -name "*.go")
PACKAGES = $$(go list ./... )
BIN = bin/cfdyndns
VERSION = $(shell git describe --match 'v[0-9]*' --dirty='.m' --always --tags)

export GO111MODULE=on
export CGO_ENABLED=0

.PHONY: default
default: clean test build

.PHONY: all
all: clean fmt test build

.PHONY: build
build: $(BIN)


$(BIN):
	go build -v \
		-tags release \
		-ldflags="-X main.Version=${VERSION}" \
		-o $(BIN) \
		cmd/cfdyndns/cfdyndns.go

.PHONY: clean
clean:
	rm -rfv bin

.PHONY: fmt
fmt:
	gofmt -l -s -w $(FILES)

.PHONY: test
test:
	go test -cover $(PACKAGES)

.PHONY: docker-build
docker-build:
	docker build -t ${DOCKER_REGISTRY}cfdyndns:${VERSION} .
	docker tag ${DOCKER_REGISTRY}cfdyndns:${VERSION} ${DOCKER_REGISTRY}cfdyndns:latest