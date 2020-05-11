#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted=$(cd "$1" && "$script_dirS"/student/mastertheLS)
	expected=$(cd "$1" && "$script_dirS"/solutions/mastertheLS)

	diff <(echo "$submitted") <(echo "$expected")
}

challenge cl-camp1/folder1
challenge cl-camp1/folder2
