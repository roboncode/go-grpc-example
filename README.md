# Sample gRPC Service in Go

The following gRPC sample project is built in Go and connects to a MongoDB. The project also
exposes the services as HTTP. In the examples directory you will see both gRPC and HTTP client samples.

Authored by: Rob Taylor <roboncode@gmail.com>

## Getting Started

### Running service from CLI

```shell script
go mod download
go run cmd/example/main.go
```

### Running as docker container (using docker-compose)

```shell script
make docker
make compose
```

### Building an executable

```shell script
make
./bin/example
```

### Running example gRPC and HTTP clients

Be sure the service is running using one of the methods above

```shell script
go run examples/main.go
```

### Running HealthCheck

First install health checker: 
* https://github.com/grpc-ecosystem/grpc-health-probe

```shell script
~/go/bin/grpc-health-probe -addr=localhost:8080
```

### Run tests

```shell script
go test ./...
```

### References used for project

#### Go layout standards

* https://github.com/golang-standards/project-layout

#### gRPC 

* https://github.com/znly/docker-protobuf
* https://github.com/grpc-ecosystem/grpc-gateway
* https://medium.com/@lchenn/generate-grpc-and-protobuf-libraries-with-containers-c15ba4e4f3ad
* https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
* https://grpc.io/docs/quickstart/go/
* https://grpc.io/docs/tutorials/basic/go/
* https://github.com/gogo/protobuf

#### Heath Check (for Kubernetes)

* https://github.com/grpc-ecosystem/grpc-health-probe
* https://developpaper.com/k8s-and-health-examination-best-practice-of-grpc-service-health-examination/
* https://github.com/grpc/grpc/blob/master/doc/health-checking.md
* https://github.com/grpc/grpc-go/tree/master/health
