#!/bin/bash

cd "$(dirname $0)" 
cd ../cmd
go fmt ...
go mod tidy
ENVIRONMENT=DEVELOPMENT DB_CONNECTION_STRING=test.sqlite3 go run .