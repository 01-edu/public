#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

chmod 776 hard-perm/*

challenge() {
	submitted=$(bash "$script_dirS"/student/hard-perm.sh && ls -l "$1" | awk '{print $1}')
	expected=$(bash "$script_dirS"/solutions/hard-perm.sh && ls -l "$1" | awk '{print $1}')
	diff <(echo "$submitted") <(echo "$expected")
}

challenge hard-perm/
