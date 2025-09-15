## plus

### Instructions

In this exercise, you will make a script `plus.sh` that will take two arguments from the command line, add them and output the result.

### Usage

```console
$ ./plus.sh 2 3
5
$
```

### Hints

Here are some commands that you can combine to find the solution:

- `expr`: this command evaluates an expression and outputs the result.

- `command substitution - $(...)`: this command can be used to interpolate the output of running a command into a string. You can replace the ellipsis with your command.

You could use `expr` like this:

```console
$ expr 2 - 2
0
$
```

And `$(...)` like this:

```console
$ echo "This is the content of a file: $(cat file.txt)"
This is the content of a file: <content of the file>`
$
```

### References

- [expr](https://www.gnu.org/software/coreutils/manual/html_node/expr-invocation.html#expr-invocation).

- [command substitution](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Command-Substitution).
