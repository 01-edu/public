#!/bin/sh

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

cp -rT /app /jail
cd /jail
if ! test -f ${EXERCISE}_test.sh; then
	echo No test file found for the exercise : "$EXERCISE"
	exit 1
fi
bash ${EXERCISE}_test.sh
echo PASS
