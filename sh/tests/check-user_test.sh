#!/usr/bin/env bash

# set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted="./check-user.sh $@
    "
	expected="./check-user.sh $@
    "
	submitted+=$(bash 2>&1 "$script_dirS"/student/check-user.sh "$@")
	expected+=$(bash 2>&1 "$script_dirS"/solutions/check-user.sh "$@")

	diff -U 1 <(echo "$submitted") <(echo "$expected")
}

challenge "-i" "root"
challenge "-e" "root"

challenge "-i" "unknown_not_found"
challenge "-e" "unknown_not_found"

challenge
challenge "-i" "root" "too" "many" "args"
challenge "-t" "root"
