# Steps

### Creating proto buffer file
- create .proto file specifying the service
- make sure to mention options go_package="" to specify where the file would be generated

### Installations
- install protobuf (if not already present using brew)
- `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
- `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
- `go get -u google.golang.org/grpc`

### Generate files using
- ```protoc --go_out=api --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative example.proto```

### Implement Grpc server
- run server in a terminal (go run server/main.go) 
### Implement Grpc client
- run client to hit the api in a different terminal(go run client.go)

