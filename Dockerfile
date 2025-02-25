FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o /exporter pkg/*

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    iproute2 \
    vim-nox \
RUN mkdir -p /app
COPY --from=builder /exporter /app/exporter

CMD [ "/app/exporter", "-port", "2112", "-appid", "PEERINGMON-DEV" ]
