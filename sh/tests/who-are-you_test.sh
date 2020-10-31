#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

echo insecure >> ~/.curlrc
caddy start &>/dev/null

submitted=$(./student/who-are-you.sh)
expected=$(./solutions/who-are-you.sh)

caddy stop &>/dev/null

diff <(echo "$submitted") <(echo "$expected")
