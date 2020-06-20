#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
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
	chmod +x "$first_file"
fi

cd
cp -rT /app .

if ! test -f "${EXERCISE}_test.sh"; then
	echo No test file found for the exercise : "$EXERCISE"
	exit 1
fi

bash "${EXERCISE}_test.sh"
