.PHONY: init
init:
	go mod init github.com/kajanan1212/simple-booking-cli

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: check-format
check-format:
	gofmt -d .

.PHONY: check-lint
check-lint:
	golangci-lint run ./...

.PHONY: format
format:
	go fmt ./...

.PHONY: build
build:
	go build .

.PHONY: run
run:
	go run .
