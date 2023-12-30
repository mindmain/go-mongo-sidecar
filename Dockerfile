FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main ./cli

# Path: /app/main
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]
