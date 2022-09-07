.PHONY: build
build:
	go build -o ./bin/wb-internship-l0 ./main.go

.PHONY: test
test:
	make build
	go test -v -race -timeout 30s ./...

.PHONY: start
start:
	make build
	./bin/wb-internship-l0

.DEFAULT_GOAL := start
