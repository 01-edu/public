## better-cat

### Instructions

Create a script `better-cat.sh` that will behave similarly to the `cat` command, but with additional functionality.

The script should accept any number of filenames as arguments, and it should print the content of each file to the standard output. If no arguments are provided, the script should print the content of all files in the current directory. Before each line it should print the line number followed by a colon(`:`) and a space as in the usage examples.

The script should also support the following options:

- `-c`: If this option is provided, the script should exclude all lines that begin with the # character.
- `-l`: If this option is provided, the script should print the length of the line before each line of output.
- `-r`: If this option is provided, the script should print a summary line at the end of the output, indicating the total number of lines in the file and the count of characters in the lines printed.

### Usage

```console
$ echo "This is an example line of text." > example.txt
$ bash better-cat.sh example.txt
1: This is an example line of text.
$ bash better-cat.sh file1.txt file2.txt
Here would be the content of file1
Here would be the content of file2
$ bash better-cat.sh
Here would be the content of file1
Here would be the content of file2
Here would be the content of file3
$ echo "# This is a comment." > example.txt
$ echo "This is the second line of text." >> example.txt
$ bash better-cat.sh -c example.txt
2: This is the second line of text.
$ bash better-cat.sh -l example.txt
1 (20): # This is a comment.
2 (32): This is the second line of text.
$ bash better-cat.sh -r example.txt
1: # This is a comment.
2: This is the second line of text.
Total: 2 lines, 52 characters
$ bash better-cat.sh -cr example.txt
2: This is the second line of text.
Total: 2 lines, 32 characters
```
