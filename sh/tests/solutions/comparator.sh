#!/usr/bin/env bash

if [ "$#" -ne 2 ]; then
	echo "Error: The script only works with two arguments!"
elif ! [[ $1 =~ ^-?[0-9]*\.?[0-9]+$ ]] || ! [[ $2 =~ ^-?[0-9]*\.?[0-9]+$ ]]; then
	echo "Error: Only two numeric arguments are acceptable!"
elif [ "$1" -gt "$2" ]; then
	echo "$1 > $2"
elif [ "$1" -lt "$2" ]; then
	echo "$1 < $2"
else
	echo "$1 = $2"
fi
