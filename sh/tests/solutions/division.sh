#!/usr/bin/env bash

# Check if two arguments were provided
if [ $# -ne 2 ]
then
    echo "Error: two numbers must be provided"
    exit 0
fi

# # Check if the arguments are numeric
if ! [[ $1 =~ ^-?[0-9]*\.?[0-9]+$ ]] || ! [[ $2 =~ ^-?[0-9]*\.?[0-9]+$ ]]
then
    echo "Error: both arguments must be numeric"
    exit 0
fi

# Check if the second argument is not 0
if [ $(echo "$2 == 0" | bc) -eq 1 ]
then
    echo "Error: division by zero is not allowed"
    exit 0
fi

# Divide the first argument by the second using bc
result=$(echo "$1 / $2" | bc )

# Output the result
echo $result
