#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(./student/who-are-you.sh)
expected=$(./solutions/who-are-you.sh)

diff <(echo "$submitted") <(echo "$expected")
