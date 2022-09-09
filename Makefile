.PHONY: build
build:
	go build -o ./bin/service ./service/
	go build -o ./bin/publisher ./publisher/

.PHONY: test
test:
	make build
	go test -v -race -timeout 30s ./...

.PHONY: service
service:
	go build -o ./bin/service ./service/
	./bin/service

.PHONY: publisher
publisher:
	go build -o ./bin/publisher ./publisher/
	./bin/publisher

.DEFAULT_GOAL := service
