#!/bin/sh

set -e

if test "$CODE_EDITOR_RUN_ONLY" = true; then
	node "/jail/student/${EXERCISE}.js" "$@"
	exit
fi

node /app/test.mjs "/jail/student" "${EXERCISE}"
