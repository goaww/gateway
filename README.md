# gateway


## Install
```shell script
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```


## Gen
```shell script
protoc -I$GOPATH/protoc-3.9.1-osx-x86_64/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
   internal/gateway/echo/echoService.proto
```


```shell script
protoc -I$GOPATH/protoc-3.9.1-osx-x86_64/include  -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  internal/gateway/echo/echoService.proto
```