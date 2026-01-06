#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/in-back-ground.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

echo "- Australia is wider than the moon. The moon sits at 3400km in diameter, while Australia's diameter from east to west is almost 4000km." >facts

challenge() {
	submitted=$(cd "$1" && bash "$script_dirS"/$FILENAME)
	std_out=$(cat output.txt)
	expected=$(cd "$1" && bash "$script_dirS"/solutions/in-back-ground.sh)
	sol_out=$(cat output.txt)
	diff <(echo "$submitted") <(echo "$expected")
	diff <(echo "$std_out") <(echo "$sol_out")
}

challenge_no_output() {
	if [[ -f output.txt ]]; then
		exit 1
	fi
	submitted=$(cd "$1" && bash "$script_dirS"/$FILENAME)
	expected=$(cd "$1" && bash "$script_dirS"/solutions/in-back-ground.sh)
	diff <(echo "$submitted") <(echo "$expected")
}

if [[ $(cat $FILENAME | grep 'nohup' | wc -l) -lt 1 ]]; then
	echo "The file does not contain the required commands"
	exit 1
fi

challenge .
rm output.txt

echo "The sun is a star!" >facts
challenge_no_output .
