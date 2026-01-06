#!/usr/bin/env bash

set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted=$(bash "$script_dirS"/student/plus.sh $1 $2)
	expected=$(bash "$script_dirS"/solutions/plus.sh $1 $2)

	diff <(echo "$submitted") <(echo "$expected")
}

challenge "1" "2"
challenge "4" "2"
challenge "3" "-3"
challenge "0" "0"
challenge "2" "-3"
challenge "-4" "1"
challenge "-2" "-1"
challenge "10" "-5"
challenge "-11" "20"
