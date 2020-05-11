#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

cd cl-camp8

submitted=$(../student/skip.sh)
expected=$(../solutions/skip.sh)

diff <(echo "$submitted") <(echo "$expected")
