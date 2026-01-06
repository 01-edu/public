#!/usr/bin/env bash

set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

# check that test command was not used
if grep -q "test" "$script_dirS"/student/file-checker.sh; then
	echo "The 'test' command is not allowed in this exercise"
	exit 1
fi

if test ! -e file-checker; then
	mkdir file-checker
	cd file-checker
	touch readable-and-writable readable-only readable-and-executable readable-writable-executable
	chmod -x "$script_dirS/file-checker/readable-and-writable"
	chmod -xw "$script_dirS/file-checker/readable-only"
	chmod -w "$script_dirS/file-checker/readable-and-executable"
	chmod +x "$script_dirS/file-checker/readable-and-executable"
	chmod +x "$script_dirS/file-checker/readable-writable-executable"
	cd ..
fi

challenge() {
	submitted=$(bash "$script_dirS"/student/file-checker.sh "$1")
	expected=$(bash "$script_dirS"/solutions/file-checker.sh "$1")

	diff <(echo "$submitted") <(echo "$expected")
}

challenge "$script_dirS/file-checker/readable-only"
challenge "$script_dirS/file-checker/readable-and-writable"
challenge "$script_dirS/file-checker/readable-and-executable"
challenge "$script_dirS/file-checker/readable-writable-executable"
rm -rf file-checker
