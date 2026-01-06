#!/usr/bin/env bash

# Check if the script was given an argument
if [ $# -ne 1 ]; then
	echo "Error: a file must be provided as an argument"
	exit 1
fi

# Check if file exists
if [ ! -e "$1" ]; then
	echo "File does not exist"
else
	echo "File exists"
fi

# Check file's permissions
if [ -r "$1" ]; then
	echo "File is readable"
else
	echo "File is not readable"
fi

if [ -w "$1" ]; then
	echo "File is writable"
else
	echo "File is not writable"
fi

if [ -x "$1" ]; then
	echo "File is executable"
else
	echo "File is not executable"
fi
