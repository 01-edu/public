#!/bin/sh

set -o errexit
set -o pipefail
IFS='
'

cd -P "$(dirname "$0")"

rm -rf student
find . -type f -executable ! -name '*.sh' -delete
