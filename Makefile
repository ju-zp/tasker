linuxbuild:                       
	rm -rf bin                        
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app ./cmd/tasker-server/main.go