#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

set -x

GOPROXY=https://goproxy.io,direct go build -o myssh github.com/lisr/myssh/cmd/myssh

# ./myssh kube

echo done 😄
