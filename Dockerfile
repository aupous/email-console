FROM golang:1.17-alpine as builder
WORKDIR /build
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app cmd/main.go