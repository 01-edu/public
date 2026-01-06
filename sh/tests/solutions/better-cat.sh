#!/bin/bash

while getopts "clr" opt; do
	case $opt in
	c) exclude_comments=true ;;
	l) show_length=true ;;
	r) show_recap=true ;;
	*)
		echo "Usage: $0 [-c] [-l] [-r] [file1] [file2] ..." >&2
		exit 1
		;;
	esac
done

shift $((OPTIND - 1))

if [[ $# -eq 0 ]]; then
	set -- *
fi

line_count=0
char_count=0

for file in "$@"; do
	if [[ ! -e $file ]]; then
		echo "$0: $file: No such file or directory" >&2
		continue
	fi

	line_num=0
	while read -r line; do
		line_num=$((line_num + 1))
		if [[ $exclude_comments && $line == \#* ]]; then
			continue
		fi
		if [[ $show_length ]]; then
			echo "$line_num ($((${#line}))): $line"
		else
			echo "$line_num: $line"
		fi
		char_count=$((char_count + ${#line}))
	done <"$file"

	line_count=$((line_count + line_num))
done

if [[ $show_recap ]]; then
	echo "Total: $line_count lines, $char_count characters"
fi

if [[ $line_count -eq 0 ]]; then
	exit 1
fi
