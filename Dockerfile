FROM golang:1.19 as build
WORKDIR /go/src/

env GO111MODULE=on
env CGO_ENABLED=0

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make clean build

FROM gcr.io/distroless/base:nonroot
USER 65532:65532
COPY --from=build /go/src/bin/cfdyndns /cfdyndns
CMD ["/cfdyndns"]
