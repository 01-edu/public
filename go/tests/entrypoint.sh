#!/bin/sh

set -o errexit
set -o pipefail
set -o nounset
IFS='
'

mkdir src/student
cd src/student

if test "$REPOSITORY"; then
	password=$(cat)
	if ! git clone --depth=1 --shallow-submodules http://root:"${password}"@"$REPOSITORY" . 2>/dev/null; then
		echo Could not clone your repository
		exit 1
	fi
else
	first_file=$(echo "$EXPECTED_FILES" | cut -d' ' -f1)
	mkdir "$(dirname $first_file)"
	cat > "$first_file"
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

# Compile and run test
cd
GOPATH=$HOME:$GOPATH
if command -v "$EXERCISE"_test &>/dev/null; then
	# The exercise is a program
	go build "student/$EXERCISE"
	"$EXERCISE"_test
else
	# The exercise is a function
	go run "$EXERCISE"_test
fi
