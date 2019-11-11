#!/bin/bash

mkdir -p ./generated
docker pull thethingsindustries/protoc
docker run --rm -v $(pwd):$(pwd) -w $(pwd) thethingsindustries/protoc \
  --gogo_out=plugins=grpc:generated \
  --grpc-gateway_out=logtostderr=true:generated \
  --validate_out="lang=gogo:generated" \
  --swagger_out=logtostderr=true:generated \
  -Iprotos ./protos/*.proto

go fmt ./generated/...
