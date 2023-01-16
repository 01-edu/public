#!/usr/bin/env bash

set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
    $(bash "$script_dirS"/student/input-redirection.sh)
    submitted=$(bash "$script_dirS"/show_info.sh)
    rm show_info.sh
    $(bash "$script_dirS"/solutions/input-redirection.sh)
    expected=$(bash "$script_dirS"/show_info.sh)
    rm show_info.sh

    diff <(echo "$submitted") <(echo "$expected")
}

challenge
