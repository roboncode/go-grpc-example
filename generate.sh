#!/bin/bash

mkdir -p ./generated
docker pull thethingsindustries/protoc
docker run --rm -v $(pwd):$(pwd) -w $(pwd) thethingsindustries/protoc \
  --gogo_out=plugins=grpc:generated \
  --grpc-gateway_out=logtostderr=true:generated \
  --validate_out="lang=go:generated" \
  --swagger_out=logtostderr=true:generated \
  -Iproto ./proto/*.proto

go fmt ./generated/...
