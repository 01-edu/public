#!/usr/bin/env bash

set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

submitted=$(bash "$script_dirS"/student/file-checker.sh "$script_dirS"/student/hello-devops.sh)
expected=$(bash "$script_dirS"/solutions/file-checker.sh "$script_dirS"/student/hello-devops.sh)

# check that test command was not used

if grep -q "test" "$script_dirS"/student/file-checker.sh
then
    echo "The 'test' command is not allowed in this exercise"
    exit 1
fi

diff <(echo "$submitted") <(echo "$expected")
