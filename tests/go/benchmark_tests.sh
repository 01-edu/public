#!/usr/bin/env bash

set -euo pipefail
IFS='
'

cd -P "$(dirname "$BASH_SOURCE")"

rm -rf student
cp -a solutions student
go test -v -json|jq -c 'select(.Action == "pass") | {Test, Elapsed}' | jq -sr 'sort_by(.Elapsed) | .[-30:] | .[] | [.Elapsed, .Test] | @tsv'
