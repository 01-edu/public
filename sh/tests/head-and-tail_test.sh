#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

echo insecure >> ~/.curlrc
caddy start &>/dev/null

submitted=$(bash student/head-and-tail.sh)
expected=$(bash solutions/head-and-tail.sh)

caddy stop &>/dev/null

diff <(echo "$submitted") <(echo "$expected")
