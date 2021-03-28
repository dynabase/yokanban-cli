# yokanban cli

## Prerequisites

- [Go](https://golang.org/doc/install)
- [Cobra library](https://github.com/spf13/cobra#readme)
- [Cobra generator](https://github.com/spf13/cobra/blob/master/cobra/README.md)
- [golangci-lint](https://golangci-lint.run/usage/install/)


To set up all prerequisites ensure that you have a derivative of [make](https://www.gnu.org/software/make/) installed on your system and run

    make deps-dev

This will take care to install all dependencies needed for development.

## Development

Check the guidelines for

- Conventions: https://golang.org/doc/effective_go
- Directory structure: [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

### Install modules

    make deps

### Run your application

    go run main.go
    go run main.go <command> <args> <flags>

### Add new CLI commands

See: https://github.com/spf13/cobra/blob/master/cobra/README.md#cobra-add

    cobra add <command>

### Linting

We are using `golangci-lint` as linter framework that aggregates several linters at once.
Please consider [integrating](https://golangci-lint.run/usage/integrations/) it as a linter in your editor of choice.

Before creating a PR, please ensure your changes follow our linting rules

    make lint
