# syntax=docker/dockerfile:1

FROM golang:1.17-alpine as builder
LABEL maintainer="Erick Amorim <github.com/ericklima-ca>"
ENV CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN go build -o /mailmango

FROM gcr.io/distroless/base-debian11
COPY --from=builder /mailmango /mailmango
ENTRYPOINT ["/mailmango"]