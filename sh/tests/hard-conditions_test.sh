#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

echo "echo hello word" >test-hard-conditions.sh
echo "echo hello word" >non-ex
chmod +x test-hard-conditions.sh

submitted=$(bash "$script_dirS"/student/hard-conditions.sh test-hard-conditions.sh)
expected=$(bash "$script_dirS"/solutions/hard-conditions.sh test-hard-conditions.sh)
diff <(echo "$submitted") <(echo "$expected")

submitted=$(bash "$script_dirS"/student/hard-conditions.sh non-ex)
expected=$(bash "$script_dirS"/solutions/hard-conditions.sh non-ex)
diff <(echo "$submitted") <(echo "$expected")
