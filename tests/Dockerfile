FROM golang:1.13-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /go/src/tasker/tests

CMD CGO_ENABLED=0 go test 

