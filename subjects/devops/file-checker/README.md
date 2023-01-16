## file-checker

### Instructions

In this exercise you will make a script `file-checker.sh` that will check the status of a given file using if statements.

Your script should make the following checks:

- File exists.

- File is readable.

- File is writable.

- File is executable.

It should output the status as in the example in the usage.

You should also handle what to do when no file is provided. The use of `test` command is not allowed for this exercise.

### Usage

```console
$ ./file-checker.sh file.txt
File exists
File is readable
File is not writable
File is not executable
$ ls -l file.txt
-r--r--r-- 1 user user 0 Jan 12 08:26 file.txt
$ ./file-checker.sh
Error: No file provided
$
```

### Hints

[bash conditional expressions](https://www.gnu.org/software/bash/manual/html_node/Bash-Conditional-Expressions.html) can be used to solve this exercise. When using those expressions, you only need to take into account it's default behavior.

For example for `script.sh`:

```bash
#!/usr/bin/env bash

if [ -s $1] ; then
    echo "File size greater than 0"
fi
```

And its output:

```console
$ ./script.sh file.txt
File size greater than 0
$
```
