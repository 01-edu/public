## File struct

### Instructions

Create the files and directories so that when you use the command `tree` below, the output will look like this:

```console
$ tree struct/
struct/
├── 0
├── 1
├── 2
├── 3
│   └── text.txt
├── 4
│   └── text2.txt
├── 5
├── 6
├── 7
├── 8
├── 9
└── A
    └── text3.txt
```

Once it is done, use the command below to create the file `file-struct.tar` to be submitted.

```console
$ tar -cf file-struct.tar *
$ ls
0  1  2  3  4  5  6  7  8  9  A  file-struct.tar
```

**Only `file-struct.tar` should be submitted.**

### Hints

In order to create new directories or folders in Linux or Unix-like operating systems you need to use the `mkdir` command. `mkdir` stands for “make directory.” you can use it like so:

```console
$ mkdir my_folder
$ ls
my_folder
```

In this exercise you wil use the command `tree` to see the file structure as in the example below. `tree` is a recursive directory listing program that produces a depth-indented listing of files. With no arguments, `tree` lists the files in the current directory.

You will also need to use the `tar` command which helps to create, extract, and list archive contents. You can find more about the command in the link below.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!

### References

- [Tar](https://www.gnu.org/software/tar)
