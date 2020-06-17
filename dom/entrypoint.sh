#!/bin/sh

set -o noglob
set -o errexit
set -o nounset
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
fi

cd

node --unhandled-rejections=strict /app/test.js "${EXERCISE}"
