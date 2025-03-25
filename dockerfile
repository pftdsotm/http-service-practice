FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY . .              
RUN go build -o app ./cmd/server
