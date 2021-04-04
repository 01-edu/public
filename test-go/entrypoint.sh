#!/bin/sh

set -e

cp -r /public .
cp -a student piscine-go
cd piscine-go

if test "$EXAM_MODE"; then
	go mod init main 2>/dev/null
	GOSUMDB=off go get github.com/01-edu/z01@v0.1.0 2>/dev/null
fi

if test "$EXAM_RUN_ONLY" = true; then
	go build -o exe "./$EXERCISE"
	./exe "$@"
	exit
fi

if ! test "$EXAM_MODE"; then
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

cd ~/public/go/tests

# Compile and run test
if command -v "${EXERCISE}_test" >/dev/null 2>&1; then
	# The exercise is a program
	"${EXERCISE}_test"
else
	# The exercise is a function
	go run "./tests/${EXERCISE}_test"
fi
