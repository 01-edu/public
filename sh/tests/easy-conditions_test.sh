#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

for i in $(seq 1 5); do
	export X=$(shuf -i 1-20 -n 1)
	export Y=$(shuf -i 1-30 -n 1)
	submitted=$(bash "$script_dirS"/student/easy-conditions.sh)
	expected=$(bash "$script_dirS"/solutions/easy-conditions.sh)
	diff <(echo "$submitted") <(echo "$expected")
done
