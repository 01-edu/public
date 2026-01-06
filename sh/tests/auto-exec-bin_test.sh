#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/auto-exec-bin.sh"
BINFILE="~/myBins/01exec"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

setupbin() {
	if [ -f ${BINFILE} ]; then
		echo "bin already exists!"
	else
		mkdir -p ~/myBins
		echo "echo Hello 01 Scripting Pool" >$HOME/myBins/01exec
	fi
	chmod +x $HOME/myBins/01exec
}

challenge() {
	OLD_PATH=$PATH
	# run soultion script
	source "$script_dirS"/$FILENAME
	submitted=$(cd / && 01exec)

	PATH=$OLD_PATH
	# run student script
	source "$script_dirS"/solutions/auto-exec-bin.sh
	expected=$(cd / && 01exec)

	# diff
	diff <(echo "$submitted") <(echo "$expected")
}

# True if FILE exists and is a regular file
if [ -f ${FILENAME} ]; then
	# FILE exists and it's not empty
	if [ -s ${FILENAME} ]; then
		if [[ $(cat $FILENAME | grep echo | wc -l) -ne 0 ]]; then
			echo "echo is not allowed in this exercise!"
			exit 1
		fi
		setupbin
		challenge
	else
		echo "The file exist but is empty"
		exit 1
	fi
else
	echo "File does not exist"
	exit 1
fi
