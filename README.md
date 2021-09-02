# Proto.Actor GO Sample App
[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/polarbit/protoacto-go-sample)

##### How To Run
go run . worker -c 6331 -s 8080 -m 127.0.0.1:6331 -m 127.0.0.1:6332 -m 127.0.0.1:6333  
go run . worker -c 6332 -s 8081 -m 127.0.0.1:6331 -m 127.0.0.1:6332 -m 127.0.0.1:6333  
go run . worker -c 6333 -s 8082 -m 127.0.0.1:6331 -m 127.0.0.1:6332 -m 127.0.0.1:6333  
go run . client

##### Notes
- At the moment, *client* only connects to "6331" worker host.
- If that host goes down, client starts to fail.
- In order to fix that, we need to provide all hosts to client cluster config. (From command arguments) 


##### Tools To Install

```bash
$ brew install protobuf
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

go install github.com/gogo/protobuf/protoc-gen-gofast
go install  github.com/gogo/protobuf/protoc-gen-gogoslick
go install github.com/AsynkronIT/protoactor-go/protobuf/protoc-gen-gograinv2@dev
```
*Also see .gitpod and .gitpod.Dockerfile*

##### Generate go types and services
`protoc --gogoslick_out=. ./messages/protos.proto`
`protoc --gograinv2_out=. ./messages/protos.proto`

*Below is not working yet*:
`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos.proto`

##### Readings
- https://proto.actor/docs/clusterintro/
- https://proto.actor/docs/cluster-partitions/
- https://github.com/oklahomer/protoactor-go-sender-example


##### Tool Links
https://github.com/protocolbuffers/protobuf
https://github.com/protocolbuffers/protobuf-go
https://grpc.io/docs/languages/go/quickstart/
https://github.com/gogo/protobuf
