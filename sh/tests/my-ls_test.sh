#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/my-ls.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
    unalias my-ls

   	$(bash "$script_dirS"/$FILENAME)

    submitted=$(cd "$1" && my-ls)

    unalias my-ls

	$(bash "$script_dirS"/solutions/my-ls.sh)
    expected=$(cd "$1" && my-ls)

    # diff
	diff <(echo "$submitted") <(echo "$expected")
}

# True if FILE exists and is a regular file
if [ -f ${FILENAME} ]; then
    # FILE exists and it's not empty
    if [ -s ${FILENAME} ]; then
        if [[ $(cat $FILENAME | grep echo | wc -l) -ne 0 ]]; then
            echo "echo is not allowed in this exercise!"
            exit 1
        fi
        challenge my-ls/folder1
    else
        echo "The file exist but is empty"
        exit 1
    fi
else
    echo "File does not exist"
    exit 1
fi