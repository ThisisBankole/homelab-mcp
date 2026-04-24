FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o homelab-mcp .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/homelab-mcp .
EXPOSE 8081
CMD ["./homelab-mcp"]