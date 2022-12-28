## easy-conditions

### Instructions

Create a script `easy-conditions.sh` which will verify the following:

It will be given a variable "X" and "Y" and You have to check:

- If the variable "X" is grater than 5 or not. If it is greater than 5, the script will print "X is greater than 5" and if it's not greater than 5, the script will print "X is less than 5".

- If the variable "Y" is less than 20 or not. If it is less than 20, the script will print "Y is less than 20" and if it's greater than 20, the script will print "Y is greater than 20".

### Usage

```console
$ echo $X $Y
6 14
$ ./easy-conditions.sh
X is greater than 5
Y is less than 20
$ echo $X $Y
3 29
$ ./easy-conditions.sh
X is less than 5
Y is greater than 20
$
```

### Hints

The `if` condition in the shell is a control structure that allows you to execute a command or block of commands based on a specified condition. The if condition has the following syntax:

```console
if CONDITION; then
  COMMAND1
  COMMAND2
  ...
fi
```

In this syntax, CONDITION is a test or expression that returns a boolean value (true or false). If the CONDITION is true, the commands inside the then block are executed. If the CONDITION is false, the commands inside the then block are skipped.

The "-gt", "-lt", and "-eq" operators are used in the shell to perform tests and comparisons on values. These operators are commonly used with the [ command (also known as the test command) to check the value of a variable or expression.

Here is a summary of the "-gt", "-lt", and "-eq" operators:

- "-gt": Greater than. This operator checks if the value on the left is greater than the value on the right.
- "-lt": Less than. This operator checks if the value on the left is less than the value on the right.
- "-eq": Equal to. This operator checks if the value on the left is equal to the value on the right.

Here are some examples of how to use the "-gt":

```console
# Set the variables "x" and "y"
x=5
y=10

# Check if "x" is greater than "y"
if [ "$x" -gt "$y" ]; then
  echo "x is greater than y"
fi
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
