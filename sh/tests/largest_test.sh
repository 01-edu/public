#!/usr/bin/env bash

set -euo pipefail
IFS='
'

FILENAME="student/largest.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge () {
    submitted=$(cd "$1" && bash "$script_dirS"/"$FILENAME")
    expected=$(cd "$1" && bash "$script_dirS"/solutions/largest.sh)

    diff <(echo "$submitted") <(echo "$expected")
}

if [ $(cat "$script_dirS"/"$FILENAME" | grep echo) ]; then
    echo "echo is not allowed in this exercise!" 
    exit 1
fi

challenge largest-dir/folder1
challenge largest-dir/folder2
challenge largest-dir
