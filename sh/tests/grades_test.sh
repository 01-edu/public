#!/usr/bin/env bash

IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {

    declare -a equals=("Student1 90" "Student2 70" "Student3 50")
    declare -a differents=("Student1 92" "Student2 75" "Student3 55" "Student4 25")

    for e in "${equals[@]}"; do
        name=$(echo "$e" | awk '{print $1}')
        grade=$(echo "$e" | awk '{print $2}')
        submitted=$(echo -e "$name $grade" | bash -c ""$script_dirS"/student/grades.sh 1")
        expected=$(echo -e "$name $grade" | bash -c ""$script_dirS"/solutions/grades.sh 1")
        diff <(echo "$submitted") <(echo "$expected")
    done
    for d in "${differents[@]}"; do
        name=$(echo "$d" | awk '{print $1}')
        grade=$(echo "$d" | awk '{print $2}')
        submitted=$(echo -e "$name $grade" | bash -c ""$script_dirS"/student/grades.sh 1")
        expected=$(echo -e "$name $grade" | bash -c ""$script_dirS"/solutions/grades.sh 1")
        diff <(echo "$submitted") <(echo "$expected")
    done

    # Testing multiple students
    submitted=$(
        bash -c ""$script_dirS"/student/grades.sh 3" <<EOF
Bob
90
Alice
75
Eve
55
EOF
    )
    expected=$(
        bash -c ""$script_dirS"/solutions/grades.sh 3" <<EOF
Bob
90
Alice
75
Eve
55
EOF
    )
    diff <(echo "$submitted") <(echo "$expected")

    submitted=$(
        bash -c ""$script_dirS"/student/grades.sh 5" <<EOF
Bob
90
Alice
75
Eve
55
john
49
Eric
65
EOF
    )
    expected=$(
        bash -c ""$script_dirS"/solutions/grades.sh 5" <<EOF
Bob
90
Alice
75
Eve
55
john
49
Eric
65
EOF
    )
    diff <(echo "$submitted") <(echo "$expected")

    # Checking if it fails with invallid grades
    submitted=$(
        bash -c ""$script_dirS"/student/grades.sh 3" <<EOF
Bob
90
Alice
150
Eve
55
EOF
    )
    expected=$(
        bash -c ""$script_dirS"/solutions/grades.sh 3" <<EOF
Bob
90
Alice
150
Eve
55
EOF
    )
    diff <(echo "$submitted") <(echo "$expected")
}

challenge
