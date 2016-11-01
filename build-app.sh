#!/bin/bash -e

set -o errexit
set -o nounset
set -o pipefail

go fmt $(go list ./... | grep -v /vendor/)
go test ./... --cover
# go test -bench=./...
go install
