## left

### Instructions

Create a file `left.sh` that will pass the content of a file `facts` to a command that will return the number of lines in that file.

### Usage

```console
$ ./left.sh
21
```

### Hints

To pass the content of a file to a command in a shell, you can use the `<` operator followed by the name of the file. For example, to pass the content of a file named `input.txt` to the `cat` command, you can use the following command:

`cat < input.txt`

The `wc` command is a utility in Unix-like operating systems that is used to count the number of lines, words, and bytes in a file or group of files. The `-l` option tells `wc` to only print the line count for each file.

For example, if you have a file called `file.txt` and you want to count the number of lines in the file, you can use the following command:

`wc -l file.txt`

This will print the line count for `file.txt`.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
