#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted=$(bash $script_dirS/student/comparator.sh "$@")
	expected=$(bash $script_dirS/solutions/comparator.sh "$@")
	diff <(echo "$submitted") <(echo "$expected")
}

for i in $(seq 1 10); do
	n1=$(shuf -i 1-20 -n 1)
	n2=$(shuf -i 1-30 -n 1)
	challenge $n1 $n2
done

challenge "0" "0"
challenge "10" "10"
challenge "-11" "-11"
challenge "14"
challenge "-11" "-11" "4"
challenge "as" "str"
