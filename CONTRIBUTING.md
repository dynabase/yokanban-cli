# yokanban cli

## Prerequisites

- [Go](https://golang.org/doc/install)
- [Cobra library](https://github.com/spf13/cobra#readme)
- [Cobra generator](https://github.com/spf13/cobra/blob/master/cobra/README.md)

```
go get -u github.com/spf13/cobra/cobra
```
- [golangci-lint](https://golangci-lint.run/usage/install/)

We are using `golangci-lint` as linter framework that aggregates several linters at once.
Please consider [integrating](https://golangci-lint.run/usage/integrations/) it as a linter in your editor of choice. 

```
## Development

Check the guidelines for

- Conventions: https://golang.org/doc/effective_go
- Directory structure: [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

### Install modules

    go mod download

### Run your application

    go run main.go
    go run main.go <command> <args> <flags>

### Add new CLI commands

See: https://github.com/spf13/cobra/blob/master/cobra/README.md#cobra-add

    cobra add <command>

### Linting

Ensure your changes follow linting rules

    make lint
