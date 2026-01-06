#!/usr/bin/env bash

IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	args=${@:1:$#-1}
	input="${@: -1}"

	submitted=$(
		bash ./student/joker-num.sh $args <<EOF
$input
EOF
	)
	expected=$(
		bash ./solutions/joker-num.sh $args <<EOF
$input
EOF
	)
	diff <(echo "$submitted") <(echo "$expected")
	if [ $? != 0 ]; then
		exit 1
	fi
}

# Good input, win
challenge 50 "1
100
49
51
50
"

# Good input, lose
challenge 50 "10
20
30
40
41
42
"

# Bad arguments
challenge "10"

# Bad arguments
challenge 0 "10"

# Bad arguments
challenge 101 "10"

# Bad arguments
challenge -20 "10"

# Bad arguments
challenge aa "10"

# Handle bad input
challenge 78 "10
aa
  

3000
-10
0
0
40
80
79
78"
