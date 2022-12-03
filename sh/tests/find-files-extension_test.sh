#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted=$(cd "$1" && bash "$script_dirS"/student/find-files-extension.sh)
	expected=$(cd "$1" && bash "$script_dirS"/solutions/find-files-extension.sh)

	diff <(echo "$submitted") <(echo "$expected")
}

challenge find-files-extension/folder1
challenge find-files-extension/folder2
