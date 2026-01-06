#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/set-env-vars.sh"

# echo not allowed
if [ -s ${FILENAME} ]; then
	if [[ $(cat $FILENAME | grep echo | wc -l) -ne 0 ]]; then
		echo "echo is not allowed in this exercise!"
		exit 1
	fi
fi

submitted=$(bash $FILENAME)
expected=$(bash solutions/set-env-vars.sh)

diff <(echo "$submitted") <(echo "$expected")
