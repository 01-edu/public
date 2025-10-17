## hard-conditions

### Instructions

Create a script `hard-conditions.sh` which will verify if a file exists and if it is executable.
If it exists and it is executable you must print "File is executable" if it is not executable or it doesn't exist you must print "File is not an executable or does not exist".

### Usage

```console
$ ls -l
-rw-rw-r-- 1 user user    19 dez 28 14:19  docs
-rwxrwxr-x 1 user user    95 dez 29 15:48  example.sh
$ ./hard-conditions.sh path/to/example.sh
File is executable
$ ./hard-conditions.sh path/to/docs.sh
File is not an executable or does not exist
$
```

### Hints

The `test` command, is a shell builtin that is used to evaluate expressions in a shell script. It has various options for performing different types of tests, such as checking the type of a file, comparing the values of two variables, or testing the status of a command.

There are two syntaxes for using the test command.

```console
$ test EXPRESSION
$ [ EXPRESSION ]
```

In a shell script, `$1` is a special variable that refers to the first argument passed to the script. Arguments are values that are passed to the script when it is run, and they can be used to modify the behavior of the script or provide input to it.

Here is an example of a simple shell script that prints the value of `$1`:

```console
#!/bin/bash

# Print the value of $1
echo "The first argument is: $1"
```

To run this script and pass an argument to it, you can use the following command:

```console
$ ./script.sh hello
The first argument is: hello
$
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
