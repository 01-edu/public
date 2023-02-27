#!/usr/bin/env bash

# create a function to be called everytime the process exit
abort () {
    rm exp.log job.log
}

trap 'abort' EXIT


set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

SUBMITTED='student/job-regist.sh'
EXPECTED='solutions/job-regist.sh'

wait_bg_jobs () {
    while [ -n "$(jobs | grep -i running)" ]; do
        echo -n "."
        sleep 1
    done
    echo
}

# test cases
one_process () {
    sleep 2 &
    source $script_dirS/$1 
}

two_processes () {
    sleep 3 &
    sleep 4 &
    source $script_dirS/$1 
}

one_process_and_suspend () {
    sleep 5 &
    source $script_dirS/$1
    kill -STOP %1
    sleep 2
    kill -CONT %1
}
# end of test cases

challenge () {
    echo "testing $1 case"
    $1 $SUBMITTED &
    $1 $EXPECTED &
    
    wait_bg_jobs

    diff <(cat exp.log) <(cat job.log)
}

challenge one_process
challenge two_processes
challenge one_process_and_suspend
