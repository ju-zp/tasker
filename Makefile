current_dir = $(shell pwd)

swagger := docker run --rm -it -e GOPATH=$(HOME)/go:/go -v $(HOME):$(HOME) -w $(current_dir) quay.io/goswagger/swagger

build:                       
	rm -rf bin                        
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app ./cmd/tasker-server/main.go

run-local:
	go run ./cmd/tasker-server/main.go

run-dev:
	docker-compose up

swagger:
	$(swagger) generate server -t gen --exclude-main --skip-models -m ../svc/models -A tasker
