# syntax=docker/dockerfile:1

## build
FROM golang:1.18.5-buster as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /dummy

## Runner
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /dummy /dummy

ENTRYPOINT ["/dummy"]