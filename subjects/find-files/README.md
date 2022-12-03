## find-files

### Instructions

"start finding ..."

Create a file `find-files.sh`, which will look for and show, in the current directory and its sub-folders:

everything that starts with an `a` or,
all the files ending with a `z` or,


- You can use this for testing: https://assets.01-edu.org/devops-branch/find-files-example.zip

- What to use : `find`

- The output should be exactly like the example bellow but with the expected name

```console
$pwd
<..>/find-files-example
$./find-files.sh
./folder2/zzzz
./folder3/asd
./folder3/sub-folder4/abc
./folder3/sub-folder4/a_correct
./folder3/sub-folder4/aefg
./folder3/asd 2
./folder3/ahmed
./folder1/aolder_lol
$
```

### Hints

`find` command is used to search and locate the list of files and directories based on conditions you specify for files that match the arguments:

```console
$find ~/
<all files and folders in the user home>
$find ~/ \( -type f \)
<all files in the user home>
$
```

May you will need to use `pipe (|)` and `&&`.

> You have to use Man or Google to know more about commands flags, in order to solve this exercice!
> Google and Man will be your friends!