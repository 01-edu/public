#!/usr/bin/env bash

num_students=$1
declare -a students

if [ $# -ne 1 ]; then
	echo "Error: expect 1 argument only!" >/dev/stderr
	exit 1
fi

for ((i = 0; i < num_students; i++)); do
	read -p "Student Name #$((i + 1)): " name
	read -p "Student Grade #$((i + 1)): " grade
	students+=("$name $grade")

	if [ $(expr "$grade" \> 100) -eq 1 ] || ! [[ "$grade" =~ ^[0-9]+$ ]]; then
		echo "Error: The grade '$grade' is not a valid input. Only numerical grades between 0 and 100 are accepted." >/dev/stderr
		exit 1
	fi
done

for student in "${students[@]}"; do
	name=$(echo $student | awk '{print $1}')
	grade=$(echo $student | awk '{print $2}')
	if [ "$grade" -ge 90 ]; then
		echo "$name: You did an excellent job!"
	elif [ "$grade" -ge 70 ]; then
		echo "$name: You did a good job!"
	elif [ "$grade" -ge 50 ]; then
		echo "$name: You need a bit more effort!"
	else
		echo "$name: You had a poor performance!"
	fi
done
