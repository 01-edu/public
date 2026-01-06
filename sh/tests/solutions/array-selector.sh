#!/usr/bin/env bash

IFS='
'

COLORS=('red' 'blue' 'green' 'white' 'black')
if [[ ! $1 =~ ^[0-9]+$ ]] || [[ $1 -le 0 ]] || [[ $1 -gt ${#COLORS[@]} ]]; then
	echo 'Error'
else
	POS=$(($1 - 1))
	echo ${COLORS[$POS]}
fi
