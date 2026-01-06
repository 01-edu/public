#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

number='^-?[0-9]+$'

do_add() {
	echo $(($1 + $2))
}

do_sub() {
	echo $(($1 - $2))
}

do_mult() {
	echo $(($1 * $2))
}

do_divide() {
	echo $(($1 / $2))
}

if [ $# != 3 ]; then
	>&2 echo "Error: expect 3 arguments"
	exit 1
elif ! [[ $1 =~ $number && $3 =~ $number ]]; then
	>&2 echo "Error: invalid number"
	exit 4
else
	case $2 in

	"+")
		echo $(do_add $1 $3)
		;;

	"-")
		echo $(do_sub $1 $3)
		;;

	"*")
		echo $(do_mult $1 $3)
		;;

	"/")
		if [ $3 == 0 ]; then
			>&2 echo "Error: division by 0"
			exit 2
		fi
		echo $(do_divide $1 $3)
		;;

	*)
		>&2 echo "Error: invalid operator"
		exit 3
		;;

	esac

fi
