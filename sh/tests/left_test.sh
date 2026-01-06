#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

FILENAME="../student/left.sh"
cd left

if [[ $(cat $FILENAME | grep echo | wc -l) -ne 0 ]]; then
	echo "echo is not allowed in this exercise!"
	exit 1
fi

if [[ $(cat $FILENAME | grep '<' | wc -l) -lt 1 ]]; then
	echo "The file does not contain the required redirection"
	exit 1
fi

submitted=$(bash $FILENAME)
expected=$(bash ../solutions/left.sh)
diff <(echo "$submitted") <(echo "$expected")
