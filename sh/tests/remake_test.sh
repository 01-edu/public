#!/usr/bin/env bash

IFS='
'

challenge() {
	if [[ $# -eq 1 && -d "$1" ]]; then
		submitted=$(bash student/remake.sh "$1")
		expected=$(bash solutions/remake.sh "$1"-expected)
		diff <(echo $submitted) <(echo $expected)
		diff <(ls -ltr $1) <(ls -ltr $1-expected)
	else
		diff <(bash student/remake.sh "$@") <(bash solutions/remake.sh "$@")
	fi
}

mkdir remake remake-expected
challenge remake
challenge
challenge not-there
rm -rf remake remake-expected
