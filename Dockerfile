FROM golang:1.13-alpine

WORKDIR /go/src/tasker
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["tasker"]
