.PHONY: build
build:
	go build -o ./cmd/service ./service/
	go build -o ./cmd/publisher ./publisher/

.PHONY: test
test:
	make build
	go test -v -race -timeout 30s ./...

.PHONY: service
service:
	go build -o ./cmd/service ./service/
	./bin/service

.PHONY: publisher
publisher:
	go build -o ./cmd/publisher ./publisher/
	./bin/publisher

.DEFAULT_GOAL := service
