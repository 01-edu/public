#!/bin/sh

set -o errexit
set -o pipefail
IFS='
'

mkdir -p src/student
cd src/student

if test "$REPOSITORY"; then
	git clone --depth=1 --shallow-submodules "$REPOSITORY" .
fi

# Check formatting
s=$(goimports -d .)
if test "$s"; then
	echo '$ goimports -d .'
	echo "$s"
	exit 1
fi

# Check restrictions
if test "$ALLOWED_FUNCTIONS"; then
	for file in $EXPECTED_FILES; do
		rc "$file" $ALLOWED_FUNCTIONS
	done
fi

# Compile test
cd
GOPATH=$GOPATH:$HOME
if command -v "$EXERCISE"_test &>/dev/null; then
	# The exercise is a program
	go build "student/$EXERCISE"
	"$EXERCISE"_test
else
	# The exercise is a function
	go run "$EXERCISE"_test
fi
