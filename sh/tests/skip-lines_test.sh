#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(bash student/skip-lines.sh)
expected=$(bash solutions/skip-lines.sh)

diff <(echo "$submitted") <(echo "$expected")
