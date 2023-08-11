#!/usr/bin/env bash

set -e

cp -r /app .
cp -a student app
cd app

if test -f "./student/${EXERCISE}.sh"; then
	chmod +x "./student/${EXERCISE}.sh"
fi

if ! test -f "${EXERCISE}_test.sh"; then
	echo "No test file found for the exercise : $EXERCISE"
	exit 1
fi

bash "${EXERCISE}_test.sh"
