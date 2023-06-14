# gRPC calculator
 - Forked demo from [Fullstack with Santosh](https://santoshk.dev/posts/2022/grpc-for-absolute-beginners-in-go/)

## Bug fixes to original code
* added support for CLI arguments for operands
* in **main.go**: ensured output is only an Int32 value by calling _result.GetResult()_, not the full *pb.AddResponse _(ex: Result:4)_

## TODO
~~* add other mathematical operations~~
~~* call service from grpcurl? (requires TLS)~~
* add persistent logging support (to a file)
* containerize the server as a remote image on Docker Hub
* add gRPC unit tests