#!/usr/bin/env bash

IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
    args=$1
    input=$2

    submitted=$(
        bash -c ""$script_dirS"/student/grades.sh $args" <<EOF
$input
EOF
    )
    expected=$(
        bash -c ""$script_dirS"/solutions/grades.sh $args" <<EOF
$input
EOF
    )
    diff <(echo "$submitted") <(echo "$expected")

    if [ $? != 0 ]; then
        exit 1
    fi
    echo $submitted
    echo $expected
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
