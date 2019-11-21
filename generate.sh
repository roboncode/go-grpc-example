#!/bin/bash

#mkdir -p ./pkg
docker pull thethingsindustries/protoc
docker run --rm -v $(pwd):$(pwd) -w $(pwd) thethingsindustries/protoc \
  --gogo_out=plugins=grpc:pkg \
  --grpc-gateway_out=logtostderr=true:pkg \
  --validate_out="lang=gogo:pkg" \
  --swagger_out=logtostderr=true:pkg \
  -Ipkg ./pkg/*.proto

go fmt ./pkg/...
