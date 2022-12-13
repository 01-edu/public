#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
    submitted=$(ls -l "$1" && bash "$script_dirS"/student/hard-perm.sh)
    expected=$(ls -l "$1" && bash "$script_dirS"/solutions/hard-perm.sh)

    diff <(echo "$submitted") <(echo "$expected")
}

challenge hard-perm/
