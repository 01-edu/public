#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

print_content() {
	mkdir -p uncompressed
	tar -xpf change-struct.tar -C uncompressed
	tree uncompressed
}

submitted=$(cd student && print_content)
expected=$(cd solutions && print_content)
diff <(echo "$submitted") <(echo "$expected")
