#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
script_dir=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

rm -rf better-cat-files
mkdir better-cat-files
cd better-cat-files

for i in {1..3}; do
	echo "This is line $i of file2" >>"file2.txt"
done

for i in {1..10}; do
	echo "This is line $i of file3" >>"file3.txt"
	echo "This is line $i of commented" >>"commented.txt"
done

echo "This is the content of file1" >"file1.txt"
echo "# commented" >"commented.txt"

challenge() {
	submitted=$(bash "$script_dir"/student/better-cat.sh "$@")
	expected=$(bash "$script_dir"/solutions/better-cat.sh "$@")
	diff <(echo "$submitted") <(echo "$expected")
}

challenge "file1.txt"
challenge "file1.txt" "file2.txt" "file3.txt"
challenge *.txt
challenge
challenge -c "commented.txt"
challenge -l "file1.txt"
challenge -r "file1.txt"
challenge -cl "commented.txt"
challenge -clr "commented.txt"

cd ..
rm -rf better-cat-files
