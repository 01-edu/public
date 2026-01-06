#!/usr/bin/env bash

IFS='
'

if [[ "$#" -ne 1 ]]; then
	echo "Error: wrong number of arguments!"
	exit 1
fi

bash "$1" &
LOG_FILE="exp.log"

while [[ -n $(jobs | grep -iv done) ]]; do
	echo $(date +"%F %T") - $(jobs %1) >>"$LOG_FILE"
	sleep 1
done
