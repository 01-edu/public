#!/bin/sh

set -o noglob
set -o errexit
set -o nounset
IFS='
'

mkdir -p src/student
cd src/student

first_file=$(echo "$EXPECTED_FILES" | cut -d' ' -f1)
dir=$(dirname "$first_file")
mkdir -p "$dir"
cat > "$first_file"

cd

# Run program
if test "$dir" = "."; then
    go build -o exe "./src/student"
else
    go build -o exe "./src/student/$EXERCISE"
fi
./exe "$@"
