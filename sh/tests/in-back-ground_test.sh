#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/in-back-ground.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
    submitted=$(cd "$1" && bash "$script_dirS"/$FILENAME)
    expected=$(cd "$1" && bash "$script_dirS"/solutions/in-back-ground.sh)
    diff <(echo "$submitted") <(echo "$expected")
    diff sol-output.log output.log
}

# True if FILE exists and is a regular file
if [ -f ${FILENAME} ]; then
    # FILE exists and it's not empty
    if [ -s ${FILENAME} ]; then
        challenge .
    else
        echo "The file exist but is empty"
        exit 1
    fi
else
    echo "File does not exist"
    exit 1
fi
