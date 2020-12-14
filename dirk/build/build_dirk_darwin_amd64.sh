#!/bin/bash
DIR="$(cd "$(dirname "$0")" && pwd)"
PARENTDIR="$(dirname "$DIR")"

export GOOS=darwin
export GOARCH=amd64

go build -o ${DIR}/dirk_${GOOS}_${GOARCH} ${PARENTDIR}/main.go
