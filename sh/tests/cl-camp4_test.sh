#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

echo insecure >>~/.curlrc
caddy start &>/dev/null

challenge() {
	submitted=$(bash student/myfamily.sh)
	expected=$(bash solutions/myfamily.sh)

	diff <(echo "$submitted") <(echo "$expected")
}

HERO_ID=1 challenge
HERO_ID=70 challenge

caddy stop &>/dev/null
