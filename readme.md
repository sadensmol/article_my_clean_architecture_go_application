This is related to the following article:
[My Clean Architecture Go Application](https://medium.com/@sadensmol/my-clean-architecture-go-application-e4611b1754cb)

# install required dev tools

```shell 
go install \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc
```


# start it
```shell
make up
```

# run tests

```shell
make up
go test ./...
```
