FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o homelab-mcp .

FROM alpine:latest
RUN apk add --no-cache docker-cli iproute2
WORKDIR /app
COPY --from=builder /app/homelab-mcp .
EXPOSE 8082
CMD ["./homelab-mcp"]