## ztail

### Instructions

Write a program that behaves like a simplified `tail` command that takes at least one file as an argument.

The only option to be handled is `-c` and will be used in all the tests as the first argument, with positive values.

For this program the `os` package can be used.

Handle the errors by returning a non-zero exit status but process all the files.

If several files are given, print a newline and the file name between each one of them (see below).

### Usage

If `file1.txt` & `file2.txt` contains :

```
abcdefghijklmnopqrstuvwxyz

```

**Note** that the files above end with a new line.

Normal cases :

```console
$ go run . -c 4 file1.txt
xyz
$ go run . -c 4 file1.txt file2.txt
==> file1.txt <==
xyz

==> file2.txt <==
xyz
$
```

Error cases :

```console
$ go run . -c 4 file1.txt nonexisting1.txt file2.txt nonexisting2.txt
==> file1.txt <==
xyz
open nonexisting1.txt: no such file or directory

==> file2.txt <==
xyz
open nonexisting2.txt: no such file or directory
$ echo $?
1
$
```
