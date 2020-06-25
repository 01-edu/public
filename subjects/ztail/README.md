## ztail

### Instructions

Write a program that behaves like a simplified `tail` command but which takes at least one file as argument.

The only option to be handled is `-c` and will be used in all the tests as the first argument, with positive values.

For this program the `os` package can be used.

Handle the errors by returning a non-zero exit status but process all the files.

If several files are given, print a new line and the file name between each one of them (see below).

### Usage

If `file1.txt` & `file2.txt` contains :

```
abcdefghijklmnopqrstuvwxyz
```

Normal cases :

```
$ ./ztail -c 4 file1.txt
xyz
$ ./ztail -c 4 file1.txt file2.txt
==> file1.txt <==
xyz

==> file2.txt <==
xyz
$
```

Error cases :

```
$ ./ztail -c 4 file1.txt nonexisting1.txt file2.txt nonexisting2.txt
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
