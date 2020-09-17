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
	git clone --quiet --depth=1 --shallow-submodules https://root:"${password}"@"$REPOSITORY" .
else
	first_file=$(echo "$EXPECTED_FILES" | cut -d' ' -f1)
	mkdir -p "$(dirname "$first_file")"
	cat > "$first_file"
fi

set +o nounset # TODO: Remove me after this variable is always set in all/tester/main.go
if test "$SKIP_FORMATTING"; then
	s=$(goimports -d .)
	if test "$s"; then
		echo 'Your Go files are not correctly formatted :'
		echo
		echo '$ goimports -d .'
		echo "$s"
		exit 1
	fi
fi
set -o nounset # TODO: Remove me after this variable is always set in all/tester/main.go

if find . -type f -name '*.go' -exec grep -qE 'print(ln)?\(' {} +; then
	echo "Your Go files cannot use print & println builtins"
	exit 1
fi

# Check restrictions
if test "$ALLOWED_FUNCTIONS" && test "$EXPECTED_FILES"; then
	IFS=' '
	first_file=$(echo "$EXPECTED_FILES" | cut -d' ' -f1)
	# shellcheck disable=SC2086
	rc "$first_file" $ALLOWED_FUNCTIONS
fi
IFS='
'
# Compile and run test
cd
GOPATH=$HOME:$GOPATH
if command -v "${EXERCISE}_test" >/dev/null 2>&1; then
	# The exercise is a program
	go build "./src/student/$EXERCISE"
	"${EXERCISE}_test"
else
	# The exercise is a function
	go build "func/${EXERCISE}_test"
	"./${EXERCISE}_test"
fi
