# gRPC calculator
## A simple gRPC calling remote objects locally with high performance using protocol buffers and Go
* forked demo from [Fullstack with Santosh](https://santoshk.dev/posts/2022/grpc-for-absolute-beginners-in-go/)
* published to [Docker Hub](https://hub.docker.com/repository/docker/jasonsalas/grpc-calculator/) and [GitHub](https://github.com/jasonsalas/grpc-calculator)

## Bug fixes to original code
* added support for CLI arguments for operands
* in **main.go**: ensured output is only an Int32 value by calling _result.GetResult()_, not the full *pb.AddResponse (ex: _4_ vs _Result:4_)

## Usage notes
* in Docker, the port must be exposed to 50052:50052
* the server and client must be executed from the root directory for logging to work (ex: running _./bin/client_)

## TODO
* push repo to GitHub
* add gRPC unit tests
~~* add persistent logging support (to a file)~~
~~* error handling: divide by zero~~
~~* containerize the server as a remote image on Docker Hub~~
~~* call service from grpcurl? (requires TLS)~~