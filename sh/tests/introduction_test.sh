#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(bash student/hello.sh)
expected=$(bash solutions/hello.sh)

diff <(echo "$submitted") <(echo "$expected") | cat -t
