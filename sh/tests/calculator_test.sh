#!/usr/bin/env bash

# set -euo pipefail
IFS='
'

script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted="./calculator.sh $@
    "
	expected="./calculator.sh $@
    "
	submitted+=$(2>&1 bash "$script_dirS"/student/calculator.sh "$@")
	submitted+="
    exit status: $?"
	expected+=$(2>&1 bash "$script_dirS"/solutions/calculator.sh "$@")
	expected+="
    exit status: $?"

	diff -U 1 <(echo "$submitted") <(echo "$expected")
	if [ $? != 0 ]; then
		exit 1
	fi
}

# Check if student uses case statement
if [[ $(cat "$script_dirS"/student/calculator.sh | grep case | wc -l) -eq 0 ]]; then
	echo "Error: the use of case statement is mandatory"
	exit 1
fi

# Valid inputs
challenge "15" "+" "10"
challenge "15" "-" "10"
challenge "15" "/" "10"
challenge "15" "*" "10"

challenge "3491" "+" "-67"
challenge "3491" "-" "-67"
challenge "3491" "/" "-67"
challenge "3491" "*" "-67"

challenge "-3491" "+" "-67"
challenge "-3491" "-" "-67"
challenge "-3491" "/" "-67"
challenge "-3491" "*" "-67"

# Invalid inputs

challenge
challenge "-3491" "*" "-67" "10" "12"

challenge "20" "/" "0"
challenge "20" "@" "10"
challenge "10" "*" "67invalid"

# Test operators functions

source $script_dirS"/student/calculator.sh" 10 + 10 >/dev/null 2>&1

if [ $(do_add 11 14) != 25 ]; then
	echo "error in function do_add"
	exit 1
fi

if [ $(do_sub 11 14) != -3 ]; then
	echo "error in function do_sub"
	exit 1
fi

if [ $(do_mult 3 5) != 15 ]; then
	echo "error in function do_mult"
	exit 1
fi

if [ $(do_divide 50 5) != 10 ]; then
	echo "error in function do_divide"
	exit 1
fi
