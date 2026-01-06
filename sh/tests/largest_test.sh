#!/usr/bin/env bash

set -euo pipefail
IFS='
'

FILENAME="student/largest.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

rm -rf largest-dir
mkdir largest-dir
cd largest-dir

directories=(folder1 folder2/aa folder2/ba/aa folder2/aa/ba)
for dir in "${directories[@]}"; do
	mkdir -p "$dir"
done

files=(.abc 34 folder1/ab 4 folder1/ac 2 folder1/az 3 folder1/bz 3
	folder1/cz 9 folder1/.hello 21 folder1/za! 3 folder2/ab 7 folder2/ac 0
	folder2/alphabet 0 folder2/az 0 folder2/bz 4 folder2/cz 0 folder2/za! 28
	folder2/aa/aa 3 folder2/aa/az 4 folder2/aa/.salut 7 folder2/aa/ba/ab 0
	folder2/aa/ba/bz 0 folder2/ba/ac 0 folder2/ba/alphabetz 0 folder2/ba/.ola 17
	folder2/ba/aa/alphabetz! 0 folder2/ba/aa/.ciao 7 folder2/ba/aa/cz 0
	folder2/ba/aa/za! 0)

for ((i = 0; i < ${#files[@]}; i += 2)); do
	dd if=/dev/zero of="${files[i]}" bs=1 count="${files[i + 1]}" 2>/dev/null
done

cd ..

challenge() {
	submitted=$(cd "$1" && bash "$script_dirS"/"$FILENAME")
	expected=$(cd "$1" && bash "$script_dirS"/solutions/largest.sh)

	diff <(echo "$submitted") <(echo "$expected")
}

if [ $(cat "$script_dirS"/"$FILENAME" | grep echo) ]; then
	echo "echo is not allowed in this exercise!"
	exit 1
fi

challenge largest-dir/folder1
challenge largest-dir/folder2
challenge largest-dir

rm -rf largest-dir
