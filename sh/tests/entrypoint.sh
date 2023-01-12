#!/usr/bin/env bash

set -e

cp -r /app .
cp -a student app
cd app

if test -f "${EXERCISE}_test.sh"; then
	bash "${EXERCISE}_test.sh"
elif test -f "${EXERCISE}_test.py"; then
	pytest "${EXERCISE}_test.py"
else
	echo "No test file found for the exercise : $EXERCISE"
	exit 1
fi
