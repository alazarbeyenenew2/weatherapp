
FROM golang:1.22-alpine AS builder
WORKDIR /app
ADD . .
RUN go build -o bin/gateway /app/cmd/main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/gateway .
COPY --from=builder /app/config/example_config.yaml /app/config/config.yaml
EXPOSE 2000
ENTRYPOINT [ "./gateway" ]
