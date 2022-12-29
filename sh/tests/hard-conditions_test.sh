#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)
submitted=$(bash "$script_dirS"/student/hard-conditions.sh hard-conditions/test-hard-conditions.sh)
expected=$(bash "$script_dirS"/solutions/hard-conditions.sh hard-conditions/test-hard-conditions.sh)
diff <(echo "$submitted") <(echo "$expected")
submitted=$(bash "$script_dirS"/student/hard-conditions.sh hard-conditions/non-ex)
expected=$(bash "$script_dirS"/solutions/hard-conditions.sh hard-conditions/non-ex)
diff <(echo "$submitted") <(echo "$expected")
