build:                       
	rm -rf bin                        
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app ./cmd/tasker-server/main.go

run:
	docker-compose up

swagger:
	swagger generate server -t gen -f ./swagger.yml --exclude-main -A tasker
