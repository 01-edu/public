#!/bin/bash

echo "Enter directory path: "
read dir_path

# Calculate the total size of all the files
total_size=0
num_files=0
for file in $(find $dir_path -type f); do
	file_size=$(stat -c%s "$file")
	total_size=$((total_size + file_size))
	num_files=$((num_files + 1))
done

# Calculate the average size of all the files
average_size=$((total_size / num_files))

# Find the largest and smallest files
largest_file=""
largest_file_size=0
smallest_file=""
smallest_file_size=99999999999
for file in $(find $dir_path -type f); do
	file_size=$(stat -c%s "$file")
	if [ $file_size -gt $largest_file_size ]; then
		largest_file=$file
		largest_file_size=$file_size
	fi
	if [ $file_size -lt $smallest_file_size ]; then
		smallest_file=$file
		smallest_file_size=$file_size
	fi
done

# Print the results
echo "Largest file: $largest_file ($largest_file_size bytes)"
echo "Smallest file: $smallest_file ($smallest_file_size bytes)"
echo "Average file size: $average_size bytes"

# Print the files larger or equal than average size
echo "Files larger or equal to average size ($average_size bytes):"
for file in $(find $dir_path -type f); do
	file_size=$(stat -c%s "$file")
	if [ $file_size -ge $average_size ]; then
		echo "$file $file_size"
	fi
done

# Print the files smaller than average size
echo "Files smaller than average size ($average_size bytes):"
for file in $(find $dir_path -type f); do
	file_size=$(stat -c%s "$file")
	if [ $file_size -lt $average_size ]; then
		echo "$file $file_size"
	fi
done
