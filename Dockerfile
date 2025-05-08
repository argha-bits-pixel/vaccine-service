FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
COPY .env /app/.env
RUN go mod tidy
RUN go build -o ./vaccination-service cmd/main.go
 
FROM alpine:latest
WORKDIR /code
COPY --from=builder /app/vaccination-service ./vaccination-service
ENTRYPOINT ["./vaccination-service"]