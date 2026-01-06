#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

chmod 303 easy-perm/*

challenge() {
	submitted=$(bash "$script_dirS"/student/easy-perm.sh && ls -l "$1" | awk '{print $1}')
	expected=$(bash "$script_dirS"/solutions/easy-perm.sh && ls -l "$1" | awk '{print $1}')
	diff <(echo "$submitted") <(echo "$expected")
}

challenge easy-perm/
