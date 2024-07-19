#!/bin/bash

cd "$(dirname $0)" 
cd ../cmd
go fmt ...
go mod tidy
go run .