build:
	@go build -o bin/dht11_server
run: build
	@./bin/dht11_server
test:
	@go test -v ./...
