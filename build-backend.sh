#!/bin/bash

# build backend
VERSION="1.2.2"

export OLD_GOOS=$(go env GOOS)
export OLD_GOARCH=$(go env GOARCH)

mkdir bin

# buid for macos x86 64-bit
export GOOS=darwin
export GOARCH=amd64
mkdir bin/$(go env GOOS)-$(go env GOARCH)
cp .env.example ./bin/$(go env GOOS)-$(go env GOARCH)/.env.example
cp -r ./assets ./bin/$(go env GOOS)-$(go env GOARCH)
go build -o ./bin/$(go env GOOS)-$(go env GOARCH)/eleutheriapay-$VERSION-$(go env GOOS)-$(go env GOARCH)
$(go env GOPATH)/bin/rice append --exec ./bin/$(go env GOOS)-$(go env GOARCH)/eleutheriapay-$VERSION-$(go env GOOS)-$(go env GOARCH)
tar -czvf ./bin/$(go env GOOS)-$(go env GOARCH).tar.gz ./bin/$(go env GOOS)-$(go env GOARCH)

# buid for linux x86 64-bit
export GOOS=linux
export GOARCH=amd64
mkdir bin/$(go env GOOS)-$(go env GOARCH)
cp .env.example ./bin/$(go env GOOS)-$(go env GOARCH)/.env.example
cp -r ./assets ./bin/$(go env GOOS)-$(go env GOARCH)
go build -o ./bin/$(go env GOOS)-$(go env GOARCH)/eleutheriapay-$VERSION-$(go env GOOS)-$(go env GOARCH)
$(go env GOPATH)/bin/rice append --exec ./bin/$(go env GOOS)-$(go env GOARCH)/eleutheriapay-$VERSION-$(go env GOOS)-$(go env GOARCH)
tar -czvf ./bin/$(go env GOOS)-$(go env GOARCH).tar.gz ./bin/$(go env GOOS)-$(go env GOARCH)
