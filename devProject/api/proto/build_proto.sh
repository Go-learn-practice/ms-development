#!/bin/bash

# 使用 protoc 编译 project_service.proto
protoc \
  --go_out=./gen \
  --go_opt=paths=source_relative \
  --go-grpc_out=./gen \
  --go-grpc_opt=paths=source_relative \
  project_service.proto