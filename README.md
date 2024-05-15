# Introduction
1. Setup [gRPC](https://grpc.io/docs/languages/go/quickstart/)
  1.1 `go get google.golang.org/protobuf/cmd/protoc-gen-go`
  1.2 `go get google.golang.org/grpc/cmd/protoc-gen-go-grpc`
  1.3 `go get -u github.com/gofiber/fiber/v3`
  1.4 Install [protoc](https://grpc.io/docs/protoc-installation/)
2. Get the common modules
  2.1 `git submodule add git@github.com:narendrayogi/dummy-ms-common.git`
3. Run `./build.sh && ./server-go-api`