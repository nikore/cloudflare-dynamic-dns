BOT_BIN = bin/cfdyndns

export GO111MODULE=on

default: clean build

build: $(BOT_BIN)

$(BOT_BIN):
	go build -v \
	        -tags release \
        	-ldflags="-X main.version=1.1" \
	        -o $(BOT_BIN) \
	        cmd/cfdyndns/cfdyndns.go

clean:
	rm -rfv bin
