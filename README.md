# gRPC calculator
 *-* Forked demo from [Fullstack with Santosh](https://santoshk.dev/posts/2022/grpc-for-absolute-beginners-in-go/)
* published to [Docker Hub](https://hub.docker.com/repository/docker/jasonsalas/grpc-calculator/)

## Bug fixes to original code
* added support for CLI arguments for operands
* in **main.go**: ensured output is only an Int32 value by calling _result.GetResult()_, not the full *pb.AddResponse (ex: _4_ vs _Result:4_)

## TODO
* push repo to GitHub
* add persistent logging support (to a file)
* add gRPC unit tests
~~* containerize the server as a remote image on Docker Hub~~
~~* call service from grpcurl? (requires TLS)~~