.PHONY: build execute

build:
	GOOS=linux go build -o ./bin/server ./server/server.go
	GOOS=linux go build -o ./bin/client ./main.go

execute:
	./bin/server