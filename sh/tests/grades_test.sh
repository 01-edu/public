#!/usr/bin/env bash

IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {

	input="${@: -1}"
	args=${@:1:$#-1}

	bash -c ""$script_dirS"/student/grades.sh $args" 2>grades_submitted_errors 1>grades_submitted_output <<EOF
$input
EOF
	grades_submitted_status=$?

	bash -c ""$script_dirS"/solutions/grades.sh $args" 2>grades_expected_errors 1>grades_expected_output <<EOF
$input
EOF
	grades_expected_status=$?

	diff <(cat grades_submitted_output) <(cat grades_expected_output)

	if [ $? != 0 ]; then
		exit 1
	fi

	diff <(cat grades_submitted_errors) <(cat grades_expected_errors)

	if [ $? != 0 ]; then
		exit 1
	fi

	diff <(echo $grades_expected_status) <(echo $grades_submitted_status)

	if [ $? != 0 ]; then
		exit 1
	fi
}

challenge 1 "Student1
90
"
challenge 1 "Student2
70
"
challenge 1 "Student3
50
"
challenge 1 "Student4
92
"
challenge 1 "Student5
75
"
challenge 1 "Student6
55
"
challenge 1 "Student7
25
"
challenge 2 "Student1
0
Alice
100
"
challenge 3 "Bob
90
Alice
75
Eve
55
"
challenge 5 "Bob
90
Alice
75
Eve
55
john
49
Eric
65
"
challenge 3 "Bob
90
Alice
150
Eve
55
"
challenge 1 "Student1
not_good
"
challenge 1 "Student1
-75
"
challenge 1 1 "Louis
20
"
