# yokanban cli

A powerful command line interface for [yokanban](httsp://yokanban.io) written in Go.

[For contributing, please read the guidelines](CONTRIBUTING.md)

# Getting started

## Installation

Clone the repo and run

    make install

Make sure your `${GOPATH}/bin` directory is within your `$PATH` variable.
See https://golang.org/doc/gopath_code#GOPATH

Afterwards the command `yokanban` should be available. Just test it by running

    yokanban help

## Create your personal yokanban service account

- Log into https://yokanban.io
- Create a service account
- Download service account credentials as JSON file `yokanban.keys.json`
- Set environment variable

```
export YOKANBAN_API_KEYS_PATH=/path/to/your/yokanban.keys.json 
yokanban test
```

## Commands

yokanban cli commands are structured in following way `yokanban <command> <arg> <flags>`.

### Test credentials

    yokanban test

### Help

    yokanban help
