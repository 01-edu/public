#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

curl --compressed --silent --show-error --max-time 10 "https://assets.01-edu.org/superhero/all.json" |
	jq '.[] | select(.id == 70) | .name'
