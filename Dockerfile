# syntax=docker/dockerfile:1
# Builder stage
FROM golang:1.20.6-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -a -installsuffix cgo -o server ./cmd/server/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

CMD ["./server"]
