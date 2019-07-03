FROM golang:latest AS builder

RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build  -tags=jsoniter -a -o /bin/main .

FROM debian:latest
ENV ADDRESS=0x74cb34d5a97c808b02a5f56631a21c822cea1204
ENV GIN_PORT=8080
ENV ETH_ENDPOINT=http://ethereum-rpc.trustwalletapp.com
ENV GIN_MODE=release

RUN mkdir /app
COPY --from=builder /bin /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
EXPOSE 8080

CMD ["/app/main"]