FILES = $$(find . -type f -name "*.go")
PACKAGES = $$(go list ./... )
BIN = bin/cfdyndns

export GO111MODULE=on
export CGO_ENABLED=0

default: clean test build

all: clean fmt test build

build: $(BIN)

$(BIN):
	go build -v \
	        -tags release \
		-ldflags="-X main.Version=0.1" \
	        -o $(BIN) \
	        cmd/cfdyndns/cfdyndns.go

clean:
	rm -rfv bin

fmt:
	gofmt -l -s -w $(FILES)

test:
	go test -race -cover $(PACKAGES)
