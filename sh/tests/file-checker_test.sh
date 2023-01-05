#!/usr/bin/env bash

set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

submitted=$(bash "$script_dirS"/student/plus.sh $1)
expected=$(bash "$script_dirS"/solutions/plus.sh $1)

# check that test command was not used
grep -q "test" $1

if [ $? -eq 0 ]
then
    echo "The 'test' command is not allowed in this exercise"
    exit 1
fi

diff <(echo "$submitted") <(echo "$expected")
