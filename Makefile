.PHONY: build execute

build:
	GOOS=linux go build -o ./bin/server ./server/server.go
	GOOS=linux go build -o ./bin/client ./cmd/client.go

execute:
	./bin/server

compile_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator/calculator.proto

docker-run:
	docker build -t grpc-calculator .
	docker run --name protobufcalc -p 50052:50052 -d grpc-calculator
	./client -op add -op1 867 -op2 5309

docker-hub-publish:
	docker build -t jasonsalas/grpc-calculator
	docker tag grpc-calculator jasonsalas/grpc-calculator:latest
	docker push jasonsalas/grpc-calculator:latest