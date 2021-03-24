run:
	go run main.go

lint:
	golint -set_exit_status $$(go list ./...)

test:
	go test ./... -coverprofile coverage.out
	go tool cover -func coverage.out

test-dev:
	go test ./... -failfast -cover

build-release:
	env CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o build/main_linux

