#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/in-back-ground.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)
echo "- Australia is wider than the moon. The moon sits at 3400km in diameter, while Australia's diameter from east to west is almost 4000km." >facts

challenge() {
    submitted=$(cd "$1" && bash "$script_dirS"/$FILENAME)
    expected=$(cd "$1" && bash "$script_dirS"/solutions/in-back-ground.sh)
    diff <(echo "$submitted") <(echo "$expected")
}

challenge .

echo "The sun is a star!" >facts
challenge .

if [ -f sol-output ] && [ -f output ]; then
    diff sol-output output
    rm sol-output output
fi
