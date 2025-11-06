#!/usr/bin/env bash

set -euo pipefail

script_dir=$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)

challenge() {
    local submitted expected
    submitted=$(./calculator.sh "$@" 2>&1)
    expected=$(bash "$script_dir/student/calculator.sh" "$@" 2>&1)

    if ! diff -U 1 <(echo "$submitted") <(echo "$expected") &>/dev/null; then
        echo "Test failed for input: $@"
        exit 1
    fi
}

# Check if student uses case statement
if ! grep -q 'case' "$script_dir/student/calculator.sh"; then
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
source "$script_dir/student/calculator.sh" >/dev/null

do_add_test() {
    if [ $(do_add 11 14) != 25 ]; then
        echo "error in function do_add"
        exit 1
    fi
}

do_sub_test() {
    if [ $(do_sub 11 14) != -3 ]; then
        echo "error in function do_sub"
        exit 1
    fi
}

do_mult_test() {
    if [ $(do_mult 3 5) != 15 ]; then
        echo "error in function do_mult"
        exit 1
    fi
}

do_divide_test() {
    if [ $(do_divide 50 5) != 10 ]; then
        echo "error in function do_divide"
        exit 1
    fi
}

do_add_test
do_sub_test
do_mult_test
do_divide_test

echo "All tests passed successfully."
