## Change struct

### Instructions

You need to rearrange the file structure that you created earlier in the "file-struct" exercise.

Use the correct commands to create the following file structure:

    - Copy the `struct` repo that you created and change its name to `new_struct`.
    - Create the `0_to_3` and `6_to_9` folders.
    - Move the folder `0, 1, 2 and 3` inside the `0_to_3` folder.
    - Move the folder `6, 7, 8 and 9` inside the `6_to_9` folder.
    - Remove the folder `5`.
    - Rename the folder `A`, to `new_folder`
    - Copy the folder `1` inside the folder `8`

```console
$ tree new_struct/
new_struct/
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
│   │   └── 1
│   └── 9
└── new_folder
    └── text3.txt
```

Once it is done, use the command below to create the file `change-struct.tar` to be submitted.

```console
$ cd new_struct/
$ tar -cf change-struct.tar *
$ ls
0_to_3  4  6_to_9  change-struct.tar  new_folder
```

**Only `change-struct.tar` should be submitted.**

### Hints

Here is an example of how to move, rename, copy and delete a file or repo:

**_Copy_**

```console
cp `file to copy` `path of destination`
cp -r `repo to copy` `path of destination`
```

**_Move or Rename_**

```console
$ mv `file to move` `path of destination`
$ mv -r `repo to move` `path of destination`
$ ls
text.txt old_repo
$ mv text.txt new_text.txt
$ mv old_repo new_repo
$ ls
new_text.txt old_repo
```

**_Delete_**

```console
$ rm `file to remove`
$ rm -r `repo to remove`
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
