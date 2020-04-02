#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

print_content() {
	mkdir -p uncompressed
	tar -xpf done.tar -C uncompressed
	TZ=utc ls -l --time-style='+%F %R' uncompressed | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}' | sed -e 's/[[:space:]]*$//'
}

submitted=$(cd student && print_content)
expected=$(cd solutions && print_content)

diff <(echo "$submitted") <(echo "$expected")
