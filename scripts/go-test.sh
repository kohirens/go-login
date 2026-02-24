#!/bin/sh

set -e

apk --no-progress --purge add git

go mod tidy
go test -v ./...
