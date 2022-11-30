## File struct

In order to create new directories or folders in Linux or Unix-like operating systems you need to use the `mkdir` command. `mkdir` stands for “make directory.” you can use it like so:

```console
User-> mkdir my_folder
User-> ls
my_folder
```

### Instructions

Create the files and directories so that when you use the command `tree` below, the output will look like this:

```console
User -> tree struct/
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

Once it is done, use the command below to create the file `done.tar` to be submitted.

```console
User-> tar -cf done.tar *
User-> ls
0  1  2  3  4  5  6  7  8  9  A  done.tar
```

**Only `done.tar` should be submitted.**

**Tips:**

Use the command `man mkdir` to get more info.
