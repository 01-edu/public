#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted=$(cd "$1" && bash "$script_dirS"/student/teacher.sh) || (
		echo "Your script exited with a non-zero status code"
		exit 1
	)
	expected=$(cd "$1" && bash "$script_dirS"/solutions/teacher.sh)

	if ! diff -q <(echo "$submitted") <(echo "$expected") &>/dev/null; then
		echo "??? What ?? Wrong answer, teacher. Did you really do the job previously? Or was is just an \"echo\" of luck??"
		exit 1
	fi
}

MAIN_SUSPECT="Harvey Dent" challenge teacher/mystery1
MAIN_SUSPECT="The Joker" challenge teacher/mystery2
MAIN_SUSPECT="The Pinguin" challenge teacher/mystery3
