#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted=$(cd "$1" && bash "$script_dirS"/student/master-the-ls.sh)
	expected=$(cd "$1" && bash "$script_dirS"/solutions/master-the-ls.sh)

	diff <(echo "$submitted") <(echo "$expected")
}

challenge master-the-ls/folder1
challenge master-the-ls/folder2
