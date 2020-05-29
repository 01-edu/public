#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

mkdir student
cd student

if test "$REPOSITORY"; then
  password=$(cat)
  git clone --quiet --depth=1 --shallow-submodules http://root:"${password}"@"$REPOSITORY" .
else
  first_file=$(echo "$EXPECTED_FILES" | cut -d' ' -f1)
  mkdir -p "$(dirname "$first_file")"
  cat > "$first_file"
  chmod +x "$first_file"
fi

cd

node test "${EXERCISE}"
