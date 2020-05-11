#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted=$(cd "$1" && "$script_dirS"/student/lookagain.sh)
	expected=$(cd "$1" && "$script_dirS"/solutions/lookagain.sh)

	diff <(echo "$submitted") <(echo "$expected")
}

challenge cl-camp5/folder1
challenge cl-camp5/folder2
