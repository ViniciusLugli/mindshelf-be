FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOMAXPROCS=2 go build -p=2 -o server ./cmd/api

FROM alpine:3.21
WORKDIR /app

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/server ./server

EXPOSE 8080

CMD ["./server"]