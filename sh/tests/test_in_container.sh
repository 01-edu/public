#!/usr/bin/env bash

IFS='
'
cd -P "$(dirname "$0")"

NC="\033[0m"
WHT="\033[0;37m"
BLK="\033[0;30m"
RED="\033[0;31m"
YEL="\033[0;33m"
BLU="\033[0;34m"
GRN="\033[0;32m"

if [ $# != 1 ]
then
    echo "Script to execute tests for scripting piscine exercises"
    echo
    echo "!-> You must pass one argument which is the name of the exercise"
    echo "    Example: bash test_in_container.sh plus"
    echo "  > The test should be a .sh or .py file in sh/tests/"
    echo "  > The solution should be also in sh/tests/student in order to be tested"
    exit 0
fi

EXERCISE=$1

printf "  ${GRN}[Build the image for running tests]${NC}\n"
docker build -t scripting_tests .

printf "  ${GRN}[Executing tests for %s]${NC}\n" $EXERCISE
docker run --read-only \
			--network none \
			--memory 500M \
			--cpus 2.0 \
			--user 1000:1000 \
			--env EXERCISE="${EXERCISE}" \
			--env USERNAME=msessa \
			--env HOME=/jail \
			--env TMPDIR=/jail \
			--workdir /jail \
			--tmpfs /jail:size=100M,noatime,exec,nodev,nosuid,uid=1000,gid=1000,nr_inodes=5k,mode=1700 \
			--volume "$(pwd)"/student:/jail/student:ro \
			-it scripting_tests