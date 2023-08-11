#!/usr/bin/env bash

set -e

cp -r /app .
cp -a student app
cd app

if test "$CODE_EDITOR_RUN_ONLY" = true; then
	if test -f "./${EXERCISE}.sh"; then
		chmod +x "./${EXERCISE}.sh"
	fi

	# run shell programs on the code editor
	bash "./${EXERCISE}.sh" "$@"
	exit
fi

if test -f "./student/${EXERCISE}.sh"; then
	chmod +x "./student/${EXERCISE}.sh"
fi

if ! test -f "${EXERCISE}_test.sh"; then
	echo "No test file found for the exercise : $EXERCISE"
	exit 1
fi

bash "${EXERCISE}_test.sh"
