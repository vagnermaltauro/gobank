build:
	@go build -o bin/gobank

run: build
	docker-compose up -d
	sleep 5
	@./bin/gobank

test:
	@go test -v ./...
