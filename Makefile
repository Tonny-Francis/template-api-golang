.PHONY: lint
lint:
	@golangci-lint run

.PHONY: build
build:
	@go build -o bin/app cmd/app/main.go

.PHONY: run
run:
	@air

.PHONY: test
test:
	@go test -v cmd/app/main_test.go