FROM golang:latest AS builder

RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build  -tags=jsoniter -a -o /bin/main .

FROM debian:latest

RUN mkdir /app
COPY --from=builder /bin /app
COPY --from=builder /build/token-list /app/token-list
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
EXPOSE 8080

CMD ["/app/main"]