#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)
challenge() {
	touch file1.txt
	submitted=$(bash "$script_dirS"/student/details.sh && ls -l file1.txt | awk '{print $1, $2, $5, $6, $7, $8, $9}')
	expected=$(bash "$script_dirS"/solutions/details.sh && ls -l file1.txt | awk '{print $1, $2, $5, $6, $7, $8, $9}')
	diff <(echo "$submitted") <(echo "$expected")
	stat file1.txt
}

challenge
rm file1.txt
