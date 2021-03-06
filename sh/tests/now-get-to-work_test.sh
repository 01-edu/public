#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(bash student/my_answer.sh)
expected=$(bash solutions/my_answer.sh)

if ! diff -q <(echo "$submitted") <(echo "$expected") &>/dev/null; then
	echo "Wrong answer, detective."
	exit 1
fi
