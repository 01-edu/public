#!/usr/bin/env bash

# create a function to be called everytime the process exit
abort() {
	rm exp.log job.log job.sh submitted expected
}

trap 'abort' EXIT

set -euo pipefail
IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

SUBMITTED='student/job-regist.sh'
EXPECTED='solutions/job-regist.sh'

create_random_job() {
	echo "sleep $((1 + $RANDOM % 10))" >job.sh
}

wait_test_case() {
	while [ -n "$(jobs | grep -i running)" ]; do
		echo -n "."
		sleep 1
	done
	echo
}

# test cases
one_process() {
	create_random_job
	bash $script_dirS/$1 job.sh
}

wrong_numb_arguments() {
	bash $script_dirS/$1
}

# end of test cases

challenge() {
	echo "testing $1 case"
	$1 $SUBMITTED >submitted &
	$1 $EXPECTED >expected &

	wait_test_case

	diff <(cat exp.log) <(cat job.log)
	diff <(cat submitted) <(cat expected)
}

challenge one_process
challenge wrong_numb_arguments
