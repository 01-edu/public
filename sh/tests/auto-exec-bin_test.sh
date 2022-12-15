#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/auto-exec-bin.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

setupbin() {
    if [ -f ${FILENAME} ]; then
        echo "bin already exists!"
    else
        mkdir -p ~/myBins
        curl -o ~/myBins/01exec https://assets.01-edu.org/devops-branch/01exec
        chmod +x ~/myBins/01exec
        echo "bin installed!"
    fi
}

challenge() {
    # run soultion script
    export PATH=""

   	$(bash "$script_dirS"/$FILENAME)
	submitted=$(cd / && 01exec)

    # run student script
    export PATH=""

	$(bash "$script_dirS"/solutions/auto-exec-bin.sh)
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
