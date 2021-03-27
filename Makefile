all: clean build-release

clean:
	go clean
	rm -rf ./build
	rm -f coverage.out

run:
	go run main.go

lint:
	golint -set_exit_status $$(go list ./...)

test:
	go test ./... -coverprofile coverage.out
	go tool cover -func coverage.out

test-dev:
	go test ./... -failfast -cover

test-html:
	go test ./... -coverprofile coverage.out
	go tool cover -html coverage.out

build-release: clean
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o build/yokanban_linux

install:
	go build -o yokanban
	@mv yokanban ${GOPATH}/bin/
	@echo "Run 'yokanban help' for further instructions..."

.PHONY: all clean run lint test test-dev build-release install
