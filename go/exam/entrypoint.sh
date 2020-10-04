#!/bin/sh

set -e

go build -o exe ./student
./exe "$@"
