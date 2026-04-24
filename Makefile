.PHONY: build run tidy clean deploy

# Build the binary locally
build:
	go build -o homelab-mcp .

# Run locally for development
run:
	go run .

# Tidy dependencies
tidy:
	go mod tidy

# Clean built binary
clean:
	rm -f homelab-mcp

# Build Docker image
docker-build:
	docker build -t homelab-mcp .

# Deploy to arrakis via scp
deploy:
	scp -r . banks@192.168.1.39:~/homelab-mcp

# SSH into arrakis
ssh:
	ssh banks@192.168.1.39

docker-build:
	docker build -t homelab-mcp .