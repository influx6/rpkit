#!/bin/env sh

echo "Install js related binaries..."
npm install -g mocha chai

echo "Running go related tests..."
go test -v ./codecs/...
go test -v ./examples/users/tests/goside/...

echo "Running js related tests..."
cd ./examples/users/tests/jsside
mocha userservicerp.test.js
cd /go/src/github.com/gokit/rpkit

echo "Running go+js related tests..."
cd ./examples/users/tests/gojs
go run main.go &
sleep 2s
mocha userservicerp.test.js
cd /go/src/github.com/gokit/rpkit

