FROM golang:latest AS builder

RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build -a -o /bin/main .

FROM debian:latest
ENV ADDRESS=
ENV GIN_PORT=
ENV ETH_ENDPOINT=
ENV GIN_MODE=

RUN mkdir /app
COPY --from=builder /bin /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
EXPOSE 8080

CMD ["/app/main"]