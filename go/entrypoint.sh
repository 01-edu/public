#!/bin/sh

set -o errexit
set -o pipefail
IFS='
'

cp -rT /app /jail

if test "$EXPECTED_FILES"; then
	echo -n 'Formatting (with `goimports -d .`) ... '
	cd /jail/student
	s=$(goimports -d .)
	if test "$s"; then
		echo
		echo "$s"
		exit 1
	fi
	echo OK
	if test "$ALLOWED_FUNCTIONS"; then
		rc "$EXPECTED_FILES" $ALLOWED_FUNCTIONS
	fi
fi

cd /jail

if ! test -f ${EXERCISE}_test.go; then
	echo No test file found for the exercise : "$EXERCISE"
	exit 1
fi

# TODO: maybe change btree exercises so that they work without student code and then delete this line
find . -name '*_test.go' ! -name ${EXERCISE}_test.go -delete
go test -failfast '-run=(?i)'test${EXERCISE}
