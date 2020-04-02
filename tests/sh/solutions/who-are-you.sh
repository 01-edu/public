#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

r=$(curl --compressed -sSm10 https://01.alem.school/assets/superhero/all.json)

echo "$r" | jq '.[] | select(.id == 70) | .name'
