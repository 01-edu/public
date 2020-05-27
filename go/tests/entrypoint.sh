#!/bin/sh

set -o noglob
set -o errexit
set -o nounset
IFS='
'

mkdir -p src/student
cd src/student

if test "$REPOSITORY"; then
	password=$(cat)
	git clone --quiet --depth=1 --shallow-submodules http://root:"${password}"@"$REPOSITORY" .
else
	first_file=$(echo "$EXPECTED_FILES" | cut -d' ' -f1)
	mkdir -p "$(dirname "$first_file")"
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
		# shellcheck disable=SC2086
		rc "$file" $ALLOWED_FUNCTIONS
	done
fi

# Compile and run test
cd
GOPATH=$HOME:$GOPATH
if command -v "${EXERCISE}_test" >/dev/null 2>&1; then
	# The exercise is a program
	go build "./src/student/$EXERCISE"
	"${EXERCISE}_test"
else
	# The exercise is a function
	go run "func/${EXERCISE}_test"
fi
