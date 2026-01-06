#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/right.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

if test ! -e right; then
	mkdir right
	cd right
	touch sample1.txt sample2 sample3.txt sample4
	cd ..
fi

challenge() {
	$(cd "$1" && bash "$script_dirS"/$FILENAME)
	submitted=$(cat $1/filtered_files.txt)
	rm $1"/filtered_files.txt"
	$(cd "$1" && bash "$script_dirS"/solutions/right.sh)
	expected=$(cat $1/filtered_files.txt)
	diff <(echo "$submitted") <(echo "$expected")
}

challenge right
rm -r right
