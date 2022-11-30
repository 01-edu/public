## Change struct

You need to rearrange the file structure that you created earlier in the "file-struct" exercise.
Here is an example of how to move, rename, copy and delete a file or repo:

**_Copy_**

```console
cp `file to copy` `path of destination`
cp -r `repo to copy` `path of destination`
```

**_Move or Rename_**

```console
User-> mv `file to move` `path of destination`
User-> mv -r `repo to move` `path of destination`
User-> ls
text.txt old_repo
User-> mv text.txt new_text.txt
User-> mv old_repo new_repo
User-> ls
new_text.txt old_repo
```

**_Delete_**

```console
User-> rm `file to remove`
User-> rm -r `repo to remove`
```

### Instructions

Use those commands to create the following file structure:

    - Copy the `struct` repo that you created and change its name to `new_struct`.
    - Create the `0_to_3` and `6_to_9` folders.
    - Move the folder `0, 1, 2 and 3` inside the `0_to_3` folder.
    - Move the folder `6, 7, 8 and 9` inside the `6_to_9` folder.
    - Remove the folder `5`.
    - Rename the folder `10`, to `new_folder`
    - Copy the folder `1` inside the folder `8`

In this exercise you wil use the command `tree` to see the file structure as in the example bellow. `tree` is a recursive directory listing program that produces a depth-indented listing of files. With no arguments, `tree` lists the files in the current directory.

```console
User-> tree new-struct/
new-struct/
├── 0_to_3
│   ├── 0
│   ├── 1
│   ├── 2
│   └── 3
│       └── text.txt
├── 4
│   └── text2.txt
├── 6_to_9
│   ├── 6
│   ├── 7
│   ├── 8
│   └── 9
└── A
    └── text3.txt
```

Once it is done, use the command below to create the file `done.tar` to be submitted.

```console
User-> tar -cf done.tar *
User-> ls new-struct/
0_to_3  4  6_to_9  A  done.tar
```

**Only `done.tar` should be submitted.**

**Tips:**

Use the command `man <name of the command>` to get more info on some command that you need to use.
