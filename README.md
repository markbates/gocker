# Gocker

Let's you quickly run any `go *` command inside of a Docker container.

```bash
$ gocker -go-version 1.9 -go-mods off test -v ./...
```

Generates and runs a Dockerfile similar to:

```Dockerfile
FROM golang:1.9

ENV GO111MODULE=off

WORKDIR $GOPATH/src/github.com/markbates/gocker

COPY . .

RUN go get -v

RUN go test -v ./...
```

## Usage

```bash
$ gocker -h

-dry-run
    dry run
-go-mods string
    turn on/off Go Modules (default "on")
-go-version string
    go version to run (default "latest")
-path string
    path to use in Docker for WORKDIR
-with-dep
    with dep management
```
