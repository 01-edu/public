#!/bin/sh

set -o errexit
set -o pipefail
IFS='
'

cp -rT /app /jail
cd /jail
if ! test -f ${EXERCISE}_test.sh; then
	echo No test file found for the exercise : "$EXERCISE"
	exit 1
fi
sh ${EXERCISE}_test.sh
echo PASS
