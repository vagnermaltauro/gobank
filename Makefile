deps:
	@go mod download

build: deps
	@go build -o bin/gobank

run: build
	docker-compose up -d
	sleep 5
	@./bin/gobank

test: deps
	@go test -v ./...
