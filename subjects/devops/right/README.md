## right

### Instructions

Create a file `right.sh` that will get the output of a file and parse it, and then write it to a file with a specific format using a single command.

Get the output of the `ls` command, parse it with the `grep` command to filter for files without a `.txt` extension, and write the output to a file called `filtered_files.txt`, check the example bellow:

### Usage

```console
$ ls
sample1.txt sample2 sample3.txt sample4
$ ./right.sh
$ cat filtered_files.txt
sample2
sample4
$
```

### Hints

`command1 | command2 > output_file`

Here, `command1` is the command that generates the output you want to parse, and `command2` is the command that parses the output. The output of `command2` will be redirected to the file `output_file` using the `>` operator.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
