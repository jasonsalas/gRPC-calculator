# gRPC calculator
## A simple gRPC calling remote objects locally with high performance using protocol buffers and Go
* forked demo from [Fullstack with Santosh](https://santoshk.dev/posts/2022/grpc-for-absolute-beginners-in-go/)
* published to [Docker Hub](https://hub.docker.com/repository/docker/jasonsalas/grpc-calculator/) and [GitHub](https://github.com/jasonsalas/grpc-calculator)

## Bug fixes to original code
* added support for CLI arguments for operands
* **[the client app](cmd/client.go)**: ensures output is only a single Int32 value by calling _result.GetResult()_, not the full *pb.AddResponse (ex: _4_ vs _Result:4_)

## Usage notes
* ensure Docker [expicitly exposes](https://github.com/jasonsalas/gRPC-calculator/blob/master/Makefile#L15) port 50052 on the container
* the server and client must be executed from the root directory for logging to work:
    * _go run ./cmd/client.go_
    * _./bin/client_ 
* querying via [grpcurl](https://github.com/fullstorydev/grpcurl): 
    * _grpcurl -plaintext localhost:50052 list_
    * _grpcurl -plaintext -d '{"Operand1":24,"Operand2":42}' localhost:50052 calclator.Calculator/Multiply_

## TODO
* add gRPC unit tests