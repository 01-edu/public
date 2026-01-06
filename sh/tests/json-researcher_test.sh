#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/json-researcher.sh"

# True if FILE exists and is a regular file
if [ -f ${FILENAME} ]; then
	# FILE exists and it's not empty
	if [ -s ${FILENAME} ]; then
		if [[ $(cat $FILENAME | grep echo | wc -l) -ne 0 ]]; then
			echo "echo is not allowed in this exercise!"
			exit 1
		fi
		echo insecure >>~/.curlrc
		caddy start &>/dev/null

		submitted=$(bash $FILENAME)
		expected=$(bash solutions/json-researcher.sh)

		caddy stop &>/dev/null

		diff <(echo "$submitted") <(echo "$expected")
	else
		echo "The file exist but is empty"
		exit 1
	fi
else
	echo "File does not exist"
	exit 1
fi
