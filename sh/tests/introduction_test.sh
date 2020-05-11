#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(./student/hello.sh)
expected=$(./solutions/hello.sh)

diff <(echo "$submitted") <(echo "$expected")
