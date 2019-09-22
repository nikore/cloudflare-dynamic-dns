FROM alpine:3.9

RUN apk update && apk add ca-certificates
RUN mkdir -p /opt/cloudflare-dynamic-dns
COPY bin/cfdyndns /opt/cloudflare-dynamic-dns/cfdyndns
WORKDIR /opt/cloudflare-dynamic-dns
CMD ["/opt/cloudflare-dynamic-dns/cfdyndns"]
