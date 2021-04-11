version := $(shell cat VERSION | xargs)

# common build flags (also improve reproducibility)
build_flags=-ldflags='-buildid= -extldflags "-static" -X "yokanban-cli/cmd.version=$(version)"' -trimpath

.PHONY: all
all: clean deps build-release

.PHONY: deps
deps:
	go mod tidy
	go mod download

.PHONY: deps-dev
deps-dev:
	go install github.com/spf13/cobra/cobra@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go mod tidy
	go mod download

.PHONY: clean
clean:
	go clean
	rm -rf ./build
	rm -f coverage.out

.PHONY: run
run:
	go run main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

.PHONY: test
test:
	go test ./... -coverprofile coverage.out
	go tool cover -func coverage.out

.PHONY: test-dev
test-dev:
	go test ./... -failfast -cover

.PHONY: test-html
test-html:
	go test ./... -coverprofile coverage.out
	go tool cover -html coverage.out

.PHONY: build-release
build-release: clean
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a $(build_flags) -o build/yokanban_linux

.PHONY: install
install:
	go build $(build_flags) -o yokanban
	@mv yokanban ${GOPATH}/bin/
	@echo "Run 'yokanban help' for further instructions..."
