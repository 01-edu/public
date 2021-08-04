#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(bash student/explain.sh)
expected=$(bash solutions/explain.sh)

if ! diff -q <(echo "$submitted") <(echo "$expected") &>/dev/null; then
	echo "Wrong answer, detective."
	exit 1
fi
