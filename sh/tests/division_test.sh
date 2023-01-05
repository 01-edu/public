#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
    submitted=$(bash "$script_dirS"/student/division.sh $1 $2)
    expected=$(bash "$script_dirS"/solutions/division.sh $1 $2)

    diff <(echo "$submitted") <(echo "$expected")
}

challenge "10 2" "5"
challenge "4 2" "2"
challenge "0.5 0.5" "1"
challenge "0.5" "Error: two numbers must be provided"
challenge "foo bar" "Error: both arguments must be numeric"
challenge "1 0" "Error: division by zero is not allowed"
challenge "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 2" "5000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
