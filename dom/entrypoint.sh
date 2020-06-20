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
  git clone --quiet --depth=1 --shallow-submodules https://root:"${password}"@"$REPOSITORY" .
else
  first_file=$(echo "$EXPECTED_FILES" | cut -d' ' -f1)
  mkdir -p "$(dirname "$first_file")"
  cat > "$first_file"
fi

cd /app
node --no-warnings --unhandled-rejections=strict test.js "${EXERCISE}"
