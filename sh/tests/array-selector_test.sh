#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted=$(bash "$script_dirS"/student/array-selector.sh "$@")
	expected=$(bash "$script_dirS"/solutions/array-selector.sh "$@")

	diff <(echo "$submitted") <(echo "$expected")
}

# Test with numbers - out of range included
for num in {0..6}; do
	challenge $num
done

# Test with a value that is not a digit

challenge "abc"

# Test with wrong number of arguments

challenge
