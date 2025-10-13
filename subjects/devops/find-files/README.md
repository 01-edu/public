## find-files

### Instructions

"start finding ..."

Create a file `find-files.sh`, which will look for and show, in the current directory and its sub-folders:

Any entry that starts with an `a` and regular files ending with a `z`

- You can use this for testing: https://assets.01-edu.org/devops-branch/find-files-example.zip

- What to use : `find`

- The output should be exactly like the example below but with the expected name

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

You may need to use pipes `(|)` and the logical operator `&&`.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
