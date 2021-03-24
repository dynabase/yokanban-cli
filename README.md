# yokanban cli

A powerful command line interface for [yokanban](httsp://yokanban.io) written in Go.

[For contributing, please read the guidelines](CONTRIBUTING.md)

# Getting started

## Installation

    go get github.com/dynabase/yokanban-cli

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
