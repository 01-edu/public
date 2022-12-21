#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(bash student/set-internal-vars.sh)
expected=$(bash solutions/set-internal-vars.sh)

diff <(echo "$submitted") <(echo "$expected")
