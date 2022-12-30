FILES = $$(find . -type f -name "*.go")
PACKAGES = $$(go list ./... )
BIN = bin/cfdyndns
VERSION = $(shell git describe --match 'v[0-9]*' --dirty='.m' --always --tags)

export GO111MODULE=on
export CGO_ENABLED=0

.PHONY: default
default: clean test build

.PHONY: all
all: clean lint test build

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

.PHONY: lint
lint:
	go vet $(PACKAGES)
	gofmt -l -s -w $(FILES)

.PHONY: test
test:
	go test -cover $(PACKAGES)

.PHONY: docker-build
docker-build:
	docker buildx build --tag cfdyndns:${VERSION} --platform=linux/amd64,linux/arm64 .
	docker tag ${DOCKER_REGISTRY}cfdyndns:${VERSION} ${DOCKER_REGISTRY}cfdyndns:latest

.PHONY: push-image
push-image:
	@if test "$(DOCKER_REGISTRY)" = "" ; then \
        echo "DOCKER_REGISTRY but must be set in order to continue"; \
        exit 1; \
	fi
	docker tag cfdyndns:${VERSION} ${DOCKER_REGISTRY}/cfdyndns:${VERSION}
	docker tag cfdyndns:${VERSION} ${DOCKER_REGISTRY}/cfdyndns:latest
	docker push ${DOCKER_REGISTRY}/cfdyndns:${VERSION}
	docker push ${DOCKER_REGISTRY}/cfdyndns:latest
