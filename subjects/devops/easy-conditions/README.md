## easy-conditions

### Instructions

For this exercise it will be given a variable "X" and "Y" and you have to create a script `easy-conditions.sh` that will check if "X" is greater than "Y", if it is print "true" and if it's not print "false".

Expected output

```console
$ export X=6 Y=14
$ echo $X $Y
6 14
$ ./easy-conditions.sh
false
$ X=29 Y=12
$ echo $X $Y
29 12
$ ./easy-conditions.sh
true
$
```

### Hints

The `test` command, is a shell builtin that is used to evaluate expressions in a shell script. It has various options for performing different types of tests, such as checking the type of a file, comparing the values of two variables, or testing the status of a command.

There are two syntaxes for using the test command.

```console
$ test EXPRESSION
$ [ EXPRESSION ]
```

The "-gt", "-lt", and "-eq" operators are used in the shell to perform tests and comparisons on values. These operators are commonly used with the [ command (also known as the test command) to check the value of a variable or expression.

Here is a summary of the "-gt", "-lt", and "-eq" operators:

- "-gt": Greater than. This operator checks if the value on the left is greater than the value on the right.
- "-lt": Less than. This operator checks if the value on the left is less than the value on the right.
- "-eq": Equal to. This operator checks if the value on the left is equal to the value on the right.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
