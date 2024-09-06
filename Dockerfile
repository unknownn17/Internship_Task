FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY . .
RUN go build -o main cmd/main.go

CMD ["./main"]