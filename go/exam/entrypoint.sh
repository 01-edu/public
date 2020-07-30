#!/bin/sh

set -o noglob
set -o errexit
set -o nounset
IFS='
'

mkdir -p src/student
cd src/student

first_file=$(echo "$EXPECTED_FILES" | cut -d' ' -f1)
mkdir -p "$(dirname "$first_file")"
cat > "$first_file"

# Check formatting
s=$(goimports -d .)
if test "$s"; then
	echo '$ goimports -d .'
	echo "$s"
	exit 1
fi

# Check restrictions
if test "$ALLOWED_FUNCTIONS" && test "$EXPECTED_FILES"; then
	IFS=' '
	# shellcheck disable=SC2086
	rc "$first_file" $ALLOWED_FUNCTIONS
	IFS='
'
fi

cd

# Run program
go build "./src/student/$EXERCISE"
./"${EXERCISE}" "$@"
