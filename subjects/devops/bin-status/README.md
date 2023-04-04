## bin-status

### Instructions

Create the script `bin-status.sh` that will return the exit status of last command.

- Expected output:

```console
$ true ; source ./bin-status.sh
0
$ false ; source ./bin-status.sh
1
$
```

### Hints

`$?` is a variable that holds the return value of the last executed command.
`echo $?` displays 0 if the last command has been successfully executed and displays a non-zero value if some error has occurred.
The bash sets `$?` To the exit status of the last executed process. By convention 0 is a successful exit and non-zero indicates some kind of error. It can be used to check if the previous command has been executed without any errors. If it has executed successfully then it stores 0. `$?` is also useful in shell scripts as a way to decide what to do depending on how the last executed command worked by checking the exit status.

```console
$ random-binary ; echo $?
<...>
<the exit status of random-binary>
$
```

The `source` command is used in Unix-like operating systems to execute commands from a specified file in the current shell environment. When the `source` command is used, the specified file is read by the shell and executed in the same environment as the caller, without creating a new subshell.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
