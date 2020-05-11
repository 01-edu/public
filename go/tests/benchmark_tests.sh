#!/bin/sh

set -o errexit
set -o pipefail
IFS='
'

cd -P "$(dirname "$0")"

./clean.sh
cp -a solutions student
go test -v -json|jq -c 'select(.Action == "pass") | {Test, Elapsed}' | jq -sr 'sort_by(.Elapsed) | .[-30:] | .[] | [.Elapsed, .Test] | @tsv'
