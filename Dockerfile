FROM golang:1.13-alpine

RUN apk update && apk add --no-cache git

WORKDIR /go/src/github.com/ju-zp/tasker

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

EXPOSE 3000

CMD [ "bin/app"]
