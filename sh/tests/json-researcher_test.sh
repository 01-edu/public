#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

echo insecure >> ~/.curlrc
caddy start &>/dev/null

submitted=$(bash student/json-researcher.sh)
expected=$(bash solutions/json-researcher.sh)

caddy stop &>/dev/null

diff <(echo "$submitted") <(echo "$expected")
