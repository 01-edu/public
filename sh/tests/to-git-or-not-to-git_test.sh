#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(bash student/to-git-or-not-to-git.sh)
expected=$(bash solutions/to-git-or-not-to-git.sh)

diff <(echo "$submitted") <(echo "$expected")
