.PHONY: default
default: clean build-release

.PHONY: clean
clean:
	go clean
	rm -rf ./build
	rm -f coverage.out

.PHONY: run
run:
	go run main.go

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
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o build/yokanban_linux

.PHONY: install
install:
	go build -o yokanban
	@mv yokanban ${GOPATH}/bin/
	@echo "Run 'yokanban help' for further instructions..."
