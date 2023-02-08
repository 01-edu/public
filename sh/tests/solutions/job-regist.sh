#!/usr/bin/env bash

IFS='
'
PID=$(jobs -l %1 | grep -Eo '[\+\-] [0-9]+' | grep -Eo '[0-9]+')
LOG_FILE="exp.log"

while kill -0 "$PID" 2> /dev/null; do
    echo $(date +"%F %T") - $(jobs %1) >> "$LOG_FILE"
    sleep 1
done
