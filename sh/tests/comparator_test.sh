#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

# Test if there were only two arguments given

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
    if [ $# -ne 2 ]; then
        submitted=$(bash "$script_dirS"/student/comparator.sh $1)
        expected=$(bash "$script_dirS"/solutions/comparator.sh $1)
    else
        submitted=$(bash "$script_dirS"/student/comparator.sh $1 $2)
        expected=$(bash "$script_dirS"/solutions/comparator.sh $1 $2)
        diff <(echo "$submitted") <(echo "$expected")
    fi
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
