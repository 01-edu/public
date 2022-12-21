#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(bash student/set-env-vars.sh)
expected=$(bash solutions/set-env-vars.sh)

diff <(echo "$submitted") <(echo "$expected")
