
FROM golang:1.22-alpine AS builder
WORKDIR /app
ADD . .
RUN go build -o bin/auth /app/cmd/main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/auth .
COPY --from=builder /app/config/example_config.yaml /app/config/config.yaml
EXPOSE 3000
ENTRYPOINT [ "./auth" ]
