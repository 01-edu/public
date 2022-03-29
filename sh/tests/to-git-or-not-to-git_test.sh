#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
FILENAME="student/to-git-or-not-to-git.sh"

if [ -f ${FILENAME} ]; then
    if [ -s ${FILENAME} ]; then
        submitted=$(bash student/to-git-or-not-to-git.sh)
        expected=$(bash solutions/to-git-or-not-to-git.sh)
        diff <(echo "$submitted") <(echo "$expected")
    else
        echo "The file exist but empty"
        exit 1
    fi
else
    echo "File does not exist"
fi
