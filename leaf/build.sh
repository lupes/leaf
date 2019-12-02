#!/usr/bin/env bash

export GOROOT=/Users/klook/work/sdk/go1.13.4

# shellcheck disable=SC2164
cd /Users/klook/work/leaf/leaf

GOOS=linux go build -v -o leaf
