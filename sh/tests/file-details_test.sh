#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted=$(cd "$1" && bash "$script_dirS"/student/file-details.sh)
	expected=$(cd "$1" && bash "$script_dirS"/solutions/file-details.sh)

	diff <(echo "$submitted") <(echo "$expected")
}

challenge hard-perm
