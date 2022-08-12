#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/to-git-or-not-to-git.sh"
submitted=$(bash student/to-git-or-not-to-git.sh)
expected=$(bash solutions/to-git-or-not-to-git.sh)

# True if FILE exists and is a regular file
if [ -f ${FILENAME} ]; then
    # FILE exists and it's not empty
    if [ -s ${FILENAME} ]; then
        diff <(echo "$submitted") <(echo "$expected")
    else
        echo "The file exist but is empty"
    fi
else
    echo "File does not exist"
fi
