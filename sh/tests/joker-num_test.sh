test_empty_input() {
    output=$(echo -n "" | ./solutions/joker-num.sh)
    expected_output="Error: Input is empty, please try again."
    diff <(echo "$output") <(echo "$expected_output")
}

test_not_a_number() {
    output=$(echo "a" | ./solutions/joker-num.sh)
    expected_output="Error: Input is not a number, please try again."
    diff <(echo "$output") <(echo "$expected_output")
}

test_number_out_of_range() {
    output=$(echo "100001" | ./solutions/joker-num.sh)
    expected_output="Error: Number out of range, please try again."
    diff <(echo "$output") <(echo "$expected_output")
}

test_correct_guess() {
    output=$(echo "50000" | ./solutions/joker-num.sh; echo "50000" | ./solutions/joker-num.sh)
    expected_output="Congratulations! You guessed the number."
    diff <(echo "$output") <(echo "$expected_output")
}

test_guess_too_low() {
    output=$(echo "50000" | ./solutions/joker-num.sh; echo "49999" | ./solutions/joker-num.sh)
    expected_output="Go up."
    diff <(echo "$output") <(echo "$expected_output")
}

test_guess_too_high() {
    output=$(echo "50000" | ./solutions/joker-num.sh; echo "50001" | ./solutions/joker-num.sh)
    expected_output="Go down."
    diff <(echo "$output") <(echo "$expected_output")
}
test_player_one() {
    test_empty_input
    test_not_a_number
    test_number_out_of_range
}
test_player_2() {
    test_empty_input
    test_not_a_number
    test_number_out_of_range
    test_correct_guess
    test_guess_too_low
    test_guess_too_high
}

test_player_1
test_player_2