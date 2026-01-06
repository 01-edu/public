#!/bin/bash

if [[ $# != 1 ||
	-z "$1" ||
	! "$1" =~ ^[0-9]+$ ||
	"$1" -lt 1 ||
	"$1" -gt 100 ]]; then
	echo "Error: wrong argument"
	exit 1
fi

number=$1

# Start the for loop for player two
for ((tries_left = 5; tries_left > 0; tries_left--)); do
	echo "Enter your guess ($tries_left tries left):"

	read guess
	if [[ $? < 0 ]]; then
		exit 1
	fi

	if [[ -z "$guess" ||
		! "$guess" =~ ^[0-9]+$ ||
		"$guess" -lt 1 ||
		"$guess" -gt 100 ]]; then
		tries_left=$tries_left+1
		continue
	fi

	if [[ "$guess" -eq "$number" ]]; then
		echo "Congratulations, you found the number in $((5 - $tries_left + 1)) moves!"
		exit
	fi

	if [[ "$guess" -lt "$number" ]]; then
		echo "Go up"
	else
		echo "Go down"
	fi
done

echo "You lost, the number was $number"
