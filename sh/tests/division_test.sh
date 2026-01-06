#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	# Test if test command was used
	if grep -q "test" "$script_dirS"/student/division.sh; then
		echo "Error: the test command cannot be used in the student script"
		return
	fi

	submitted=$(bash "$script_dirS"/student/division.sh $@)
	expected=$(bash "$script_dirS"/solutions/division.sh $@)
	diff <(echo "$submitted") <(echo "$expected")
}

challenge "10" "2"
challenge "4" "2"
challenge "0.5" "0.5"
challenge "5" "2"
challenge "0.5"
challenge "foo" "bar"
challenge "1" "0"
challenge "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" "2"
