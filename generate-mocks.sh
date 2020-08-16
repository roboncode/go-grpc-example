#!/bin/bash

# When linking the docker port you must reference the original fs when making docker run commands.
path=$(pwd)
if [[ ${HOST_DIR} ]]; then
   path=${HOST_DIR}
fi

USE_LOCAL=${USE_LOCAL:-true}

if [ "${USE_LOCAL}" = false ]; then
  docker pull vektra/mockery
fi

doMockery() {
  MOCK_DIR=$1

  if [ "${USE_LOCAL}" = true ]; then
    mockery -dir $MOCK_DIR -output "$MOCK_DIR/mocks" -all
  else
    docker run --rm \
      -v ${path}:${path} \
      -w ${path} \
      -e GOFLAGS=-mod=vendor \
      vektra/mockery \
      -dir $MOCK_DIR -output "$MOCK_DIR/mocks" -all
  fi
}

doMockery ./generated
#doMockery ./internal/connections
#doMockery ./internal/grpc
#doMockery ./internal/grpc/interceptors
#doMockery ./internal/http
doMockery ./internal/store
