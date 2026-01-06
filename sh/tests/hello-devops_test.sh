#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/hello-devops.sh"

# True if FILE exists and is a regular file
if [ -f ${FILENAME} ]; then
	# FILE exists and it's not empty
	if [ -s ${FILENAME} ]; then
		submitted=$(bash $FILENAME)
		expected=$(bash solutions/hello-devops.sh)
		diff <(echo "$submitted") <(echo "$expected") | cat -t
	else
		echo "The file exist but is empty"
		exit 1
	fi
else
	echo "File does not exist"
	exit 1
fi
