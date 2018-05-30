#!/bin/env bash

echo "Running go related tests..."
go test -v ./users/tests/goside/...


echo "Running js related tests..."
cd ./users/tests/jsside
npm install
mocha ./userservicerp.test.js
cd /rpkit

echo "Running go+js related tests..."
cd ./users/tests/gojs
npm install
go run main.go &
sleep 2s
mocha ./userservicerp.test.js
cd /rpkit

