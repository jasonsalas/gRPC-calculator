.PHONY: build execute

build:
	GOOS=linux go build -o ./bin/server ./server/server.go
	GOOS=linux go build -o ./bin/client ./main.go

execute:
	./bin/server

compile_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator/calculator.proto