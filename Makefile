SERVICE_NAME = tasker
CURRENT_DIR = $(shell pwd)
TEST_DIR = integration/services
SERVICE_DIR = svc/services

SWAGGER_VERSION = 0.25.0

SWAGGER = /usr/local/bin/swagger-v0.25.0

build:                       
	rm -rf bin                        
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app ./main.go

run-local:
	go run ./main.go

run-dev:
	docker-compose up

swagger:
	rm -rf $(SERVICE_DIR)/$(SERVICE_NAME)
	mkdir -p $(SERVICE_DIR)/$(SERVICE_NAME)
	$(SWAGGER) generate server -t $(SERVICE_DIR)/$(SERVICE_NAME) --exclude-main -A $(SERVICE_NAME)
	rm -rf $(TEST_DIR)/$(SERVICE_NAME)
	mkdir -p $(TEST_DIR)/$(SERVICE_NAME)
	$(SWAGGER) generate client -t $(TEST_DIR)/$(SERVICE_NAME)


test:
	go test -v ./...
