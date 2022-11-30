#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

print_content() {
    mkdir -p uncompressed
    tar -xpf done.tar -C uncompressed
    tree uncompressed
}

submitted=$(cd student/struct && print_content)
expected=$(cd solutions/struct && print_content)
diff <(echo "$submitted") <(echo "$expected")
