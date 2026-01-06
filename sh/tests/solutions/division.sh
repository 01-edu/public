#!/usr/bin/env bash

# Check if two arguments were provided
if [ $# -ne 2 ]; then
	echo "Error: two numbers must be provided"

# # Check if the arguments are numeric
elif ! [[ $1 =~ ^-?[0-9]*\.?[0-9]+$ ]] || ! [[ $2 =~ ^-?[0-9]*\.?[0-9]+$ ]]; then
	echo "Error: both arguments must be numeric"

# Check if the second argument is not 0
elif [ $(echo "$2 == 0" | bc) -eq 1 ]; then
	echo "Error: division by zero is not allowed"

# Divide the first argument by the second using bc
else
	result=$(echo "$1 / $2" | bc)

	# Output the result
	echo $result
fi
