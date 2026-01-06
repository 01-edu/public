#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/head-and-tail.sh"

# True if FILE exists and is a regular file
if [ -f ${FILENAME} ]; then
	# FILE exists and it's not empty
	if [ -s ${FILENAME} ]; then
		if [[ $(cat $FILENAME | grep echo | wc -l) -ne 0 ]]; then
			echo "echo is not allowed in this exercise!"
			exit 1
		fi
		submitted=$(bash $FILENAME)
		expected=$(bash solutions/head-and-tail.sh)
		diff <(echo "$submitted") <(echo "$expected")
	else
		echo "The file exist but is empty"
		exit 1
	fi
else
	echo "File does not exist"
	exit 1
fi
