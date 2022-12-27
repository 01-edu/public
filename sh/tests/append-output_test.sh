#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

cd append-output
# reset the files.
if [ -f results.txt ]; then
    rm results.txt
    rm solution.txt
fi

# add the top two songs to the file as they are not asked.
cat songs.txt | grep "Linkin" >>results.txt

submitted=$(bash ../student/append-output.sh)
expected=$(bash ../solutions/append-output.sh)

diff <(cat results.txt) <(cat solution.txt)
