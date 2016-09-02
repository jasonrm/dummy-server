#!/bin/sh
set -o xtrace
set -o errexit
set -o pipefail

DEPS=""
BUILD_DEPS="go git build-base"
apk add --update ${DEPS} ${BUILD_DEPS}

# I wouldn't know where to start describing how much I dislike the golang build process
## Go build env
mkdir -p /go/src /go/bin
chmod -R 777 /go
export GOPATH=/go
export PATH=/go/bin:$PATH

# Install gof3r
CGO_ENABLED=0 GOOS=linux go build -o dummy-server main.go
mv dummy-server /usr/local/bin/

# Cleanup
apk del ${BUILD_DEPS}
rm -rf /var/cache/apk/*
rm -rf /go

rm /$0
