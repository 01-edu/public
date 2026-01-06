#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

print_content() {
	mkdir -p uncompressed
	tar -xpf auto-jobs.tar -C uncompressed
	cat -e uncompressed/$1
}

for i in 1 2 3 4; do
	submitted=$(cd student && print_content task$i)
	expected=$(cd solutions && print_content task$i)

	diff <(echo "$submitted") <(echo "$expected")
done
