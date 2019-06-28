#!/usr/bin/env bash

_dir=$(cd "$(dirname "$0")"; pwd)

cd $_dir

go run main.go
