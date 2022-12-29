#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)
cd "$script_dirS"
submitted=$(bash "$script_dirS"/student/hard-conditions.sh hard-conditions_test.sh)
expected=$(bash "$script_dirS"/solutions/hard-conditions.sh hard-conditions_test.sh)
diff <(echo "$submitted") <(echo "$expected")
