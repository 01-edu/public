## right

### Instructions

Create a file `right.sh` where you write a shell command that lists all the files in the current directory except those with the `.txt` extension, and save the resulting list into a file named `filtered_files.txt`. Check the example below:

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
