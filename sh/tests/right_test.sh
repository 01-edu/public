#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/right.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
    submitted=$(cd "$1" && bash "$script_dirS"/$FILENAME)
    expected=$(cd "$1" && bash "$script_dirS"/solutions/right.sh)
    diff <(echo "$submitted") <(echo "$expected")
    diff $1"/filtered_files_sol.txt" $1"/filtered_files.txt"
}

challenge right
