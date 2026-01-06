#!/usr/bin/env bash

largest=0

for i in {1..10}; do

	read -p "Enter a number: " num
	if ! [[ $num =~ ^[0-9]+$ ]]; then
		echo "ERROR: Invalid input only positive numerical characters are allowed"
		exit 1
	elif [ $num -gt $largest ]; then
		if [ $num -gt 10000 ]; then
			echo "ERROR: The number entered is too large"
			exit 1
		else
			largest=$num
		fi
	fi
done

echo "The largest number is: $largest"
