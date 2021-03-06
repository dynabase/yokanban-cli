# yokanban cli

A powerful command line interface for [yokanban](httsp://yokanban.io) written in Go.

![yokanban board](./docs/imgs/yokanban.png "yokanban board")

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

### Help

    yokanban help

### Test credentials

    yokanban test

### Board

#### Create board

    yokanban create board --title test-board

#### Delete board

    yokanban delete board --id 605f574e26f0535cfd7fd6cd

#### Get board

    yokanban get board --id 605f574e26f0535cfd7fd6cd

#### List boards

    yokanban list boards

#### Update board (title)

    yokanban update board --id 605f574e26f0535cfd7fd6cd --title test-board-udpated

### Columns

#### Create column

    yokanban create column --title test-col --board-id 605f574e26f0535cfd7fd6cd

### Cards

#### Create card

    yokanban create card --title test-card --board-id 605f574e26f0535cfd7fd6cd


# Improve your developer experience (PLANNED)

**_NOTE:_** planned - not implemented yet


## Keep context with yokanban

Develop new features without losing context due to application switches.
See how a possible development flow could look like:

```shell
# project setup
$ git clone git@github.com:MY-PROJECT.git

$ yokanban create board --title MY-PROJECT
$ yokanban create column --title "ToDo" --board MY-PROJECT
$ yokanban create column --title "In Progress" --board MY-PROJECT
$ yokanban create column --title "In Review" --board MY-PROJECT
$ yokanban create column --title "Done" --board MY-PROJECT


# prepare feature implementation
$ git checkout -b featureA

$ yokanban create card --title featureA --description "A beautiful new feature" --board MY-PROJECT --column "ToDo" --assign-me
FOO-1

# start implementing
$ yokanban move card --id "FOO-1" --column "In Progress"

$ git commit -am "[FOO-1] add beautiful function A"
$ git commit -am "[FOO-1] fix typo in beautiful function A"
$ git commit -am "[FOO-1] add documentation"

# create PR
$ git push origin featureA

Create a pull request from 'featureA' on GitHub by visiting:
	https://github.com/MY-PROJECT/pull/new/featureA

$ yokanban move card --id "FOO-1" --column "In Review"

# PR merged into main

$ yokanban move card --id "FOO-1" --column "Done"
```

## Found a bug to solve later on?

Create a ticket for it, so you won't forget it to fix it:

```shell
$ yokanban create card --color red --title "Wrong http status code" --description "Wrong http status code at REST api route PATCH /foo/bar" --board MY-PROJECT --column "ToDo"
FOO-478
```
