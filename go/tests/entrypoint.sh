#!/bin/sh

set -e

cd student

if ! test "$SKIP_FORMATTING"; then
	s=$(goimports -d .)
	if test "$s"; then
		echo 'Your Go files are not correctly formatted :'
		echo
		echo '$ goimports -d .'
		echo "$s"
		exit 1
	fi
fi

if ! find . -type f -name '*.go' | grep -q .; then
	echo "Missing Go file: $FILE"
	exit 1
fi

if find . -type f -name '*.go' -exec grep -qE 'print(ln)?\(' {} +; then
	echo "Your Go files cannot use print & println builtins"
	exit 1
fi

# Check restrictions
if test "$ALLOWED_FUNCTIONS" && test "$FILE"; then
	# shellcheck disable=SC2086
	rc "$FILE" $ALLOWED_FUNCTIONS
fi

cd

# Compile and run test
if command -v "${EXERCISE}_test" >/dev/null 2>&1; then
	# The exercise is a program
	go build -o exe "./student/$EXERCISE"
	"${EXERCISE}_test"
else
	# The exercise is a function
	mkdir src
	ln -s $(pwd)/student $(pwd)/src/student
	GOPATH=$GOPATH:$PWD
	go build "func/${EXERCISE}_test"
	"./${EXERCISE}_test"
fi
