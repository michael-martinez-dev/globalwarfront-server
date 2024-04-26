.PHONY: pre linux windows build run clean

build: pre linux windows

pre:
	@go mod tidy
	@go fmt ./...

linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/dev-server-linux ./cmd/server/main.go

windows:
	@GOOS=windows GOARCH=amd64 go build -o bin/dev-server-windows.exe ./cmd/server/main.go
	
run: pre
	@go run ./cmd/server/main.go
	
clean:
	@rm -rf ./bin/dev-*
	