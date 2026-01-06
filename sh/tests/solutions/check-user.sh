#!/usr/bin/env bash

set -euo pipefail
IFS='
'

if [ $# != 2 ]; then
	>&2 echo "Error: expect 2 arguments"
	exit 1
elif [[ $1 == "-e" ]]; then
	user_info=$(getent passwd $2)
	if [[ $user_info == "" ]]; then
		echo no
	else
		echo yes
	fi
elif [[ $1 == "-i" ]]; then
	getent passwd $2
else
	>&2 echo "Error: unknown flag"
	exit 1
fi
