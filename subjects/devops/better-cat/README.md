## better-cat

### Instructions

Create a script `better-cat.sh` that will behave similarly to the cat command, but with additional functionality.

The script should accept any number of filenames as arguments, and it should print the content of each file to the standard output. If no arguments are provided, the script should print the content of all files in the current directory.

The script should also support the following options:

- -c: If this option is provided, the script should exclude all lines that begin with the # character.
- -l: If this option is provided, the script should print the line number before each line of output.
- -r: If this option is provided, the script should print a summary line at the end of the output, indicating the total number of lines and characters printed.

If both -c and -l are provided, the line number should include only non-comment lines.

### Usage

```console
$ bash better-cat.sh file1.txt
Here would be the content of file1

$ bash better-cat.sh file1.txt file2.txt
Here would be the content of file1
Here would be the content of file2

$ bash better-cat.sh
Here would be the content of file1
Here would be the content of file2
Here would be the content of file3

$ bash better-cat.sh -c commented.txt
Here would be line 1 of file3
Here would be line 2 of file3
Here would be line 3 of file3
Here would be line 4 of file3
Here would be line 5 of file3
Here would be line 6 of file3
Here would be line 7 of file3
Here would be line 8 of file3
Here would be line 9 of file3
Here would be line 10 of file3

$ bash better-cat.sh -l file1.txt
1: Here would be the content of file1

$ bash better-cat.sh -r file1.txt
Here would be the content of file1
Total: 1 lines, 35 characters

$ bash better-cat.sh -cl commented.txt
Here would be line 1 of file3
Here would be line 2 of file3
Here would be line 3 of file3
Here would be line 4 of file3
Here would be line 5 of file3
Here would be line 6 of file3
Here would be line 7 of file3
Here would be line 8 of file3
Here would be line 9 of file3
Here would be line 10 of file3
Total: 10 lines, 301 characters

```
