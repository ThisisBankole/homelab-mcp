.PHONY: build run tidy clean deploy ssh

build:
	go build -o homelab-mcp .

run:
	go run .

tidy:
	go mod tidy

clean:
	rm -f homelab-mcp

docker-build:
	docker build -t homelab-mcp .

deploy:
	scp -r . arrakis:~/homelab-mcp

ssh:
	ssh arrakis