#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

cd cl-camp8

submitted=$(bash ../student/skip.sh)
expected=$(bash ../solutions/skip.sh)

diff <(echo "$submitted") <(echo "$expected")
