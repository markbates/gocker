FROM golang:latest

ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/markbates/gocker

COPY . .

RUN go get -v

RUN go test -v ./...
