current_dir = $(shell pwd)

swagger := docker run --rm -it -e GOPATH=$(HOME)/go:/go -v $(HOME):$(HOME) -w $(current_dir) quay.io/goswagger/swagger

build:                       
	rm -rf bin                        
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app ./main.go

run-local:
	go run ./main.go

run-dev:
	docker-compose up

swagger:
	$(swagger) generate server -t svc --exclude-main  -A tasker


test:
	docker-compose -f ./tests/docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f ./tests/docker-compose.test.yml down --volumes

test-db-up:
	docker-compose -f docker-compose.test.yml up --build db

test-db-down:
	docker-compose -f docker-compose.test.yml down --volumes db
