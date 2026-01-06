#!/usr/bin/env bash

set -euo pipefail
IFS='
'

FILENAME="student/dir-info.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

rm -rf student/dir-info-files
mkdir student/dir-info-files
cd student/dir-info-files

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

cd ../

challenge() {
	submitted=$(bash "$script_dirS"/"$FILENAME" < <(echo "$1"))
	expected=$(bash "$script_dirS"/solutions/dir-info.sh < <(echo "$1"))

	diff <(echo "$submitted") <(echo "$expected")
}

challenge "dir-info-files/folder1"
challenge "dir-info-files/folder2"
challenge "dir-info-files"

rm -rf dir-info-files
