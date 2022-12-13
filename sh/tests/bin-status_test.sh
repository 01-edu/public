#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/bin-status.sh"


challenge() {
	submitted=$(eval "$1" ; bash $FILENAME)
	expected=$(eval "$1" ; bash solutions/bin-status.sh)

	diff <(echo "$submitted") <(echo "$expected")
}


# True if FILE exists and is a regular file
if [ -f ${FILENAME} ]; then
    # FILE exists and it's not empty
    if [ -s ${FILENAME} ]; then
        challenge true
		challenge false
		challenge ls -l
		challenge ls asdasdasdasdad
    else
        echo "The file exist but is empty"
        exit 1
    fi
else
    echo "File does not exist"
    exit 1
fi