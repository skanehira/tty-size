export GO111MODULE=on
BINARY=tty-size

build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY)

docker-build: build
	docker build -t $(BINARY) .
	rm -rf $(BINARY)
