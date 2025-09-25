## greatest-of-all

### Instructions

Create a script `greatest-of-all.sh` which will ask you to input 10 numbers and then it will check what was the biggest number given. You must ask for the number using the string "Enter a number: " and then use the string "The largest number is: " to print the output like in the example below.

- Only positive numbers up to "10000" will be tested.
- If the given number is greater than "10000" you must print the error message "ERROR: The number entered is too large" and if its not a number or it is a negative number, print the error "ERROR: Invalid input only positive numerical characters are allowed". When either of these errors occurs, the script will print the error message, exit with an exit code of `1`, and will not continue to execute the next line.

### Usage

```console
$./greatest-of-all.sh
Enter a number: 1 # these numbers will be introduced by the user
Enter a number: 2
Enter a number: 3
Enter a number: 4
Enter a number: 2
Enter a number: 42
Enter a number: 1
Enter a number: 24
Enter a number: 45
Enter a number: 2
The largest number is 45
$ ./greatest-of-all.sh
Enter a number: -14
ERROR: Invalid input only positive numerical characters are allowed
./greatest-of-all.sh
Enter a number: alpha
ERROR: Invalid input only positive numerical characters are allowed
./greatest-of-all.sh
Enter a number: 10001
ERROR: The number entered is too large
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

The "-gt", "-lt", and "-eq" operators are used in the shell to perform tests and comparisons on values. These operators are commonly used with the "[" command (also known as the test command) to check the value of a variable or expression.

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

`read -p` is a command in Bash (the Unix shell) that reads a line of input from the user and stores it in a variable. The -p option allows you to specify a prompt to display to the user before reading the input.

Here is an example of how you might use read -p in a Bash script:

```console
echo "Enter your name: "
read -p 'Name: ' name
echo "Hello, $name"
```

This script will display the prompt "Enter your name: " and then wait for the user to enter their name. The name that the user enters will be stored in the name variable, and then the script will greet the user with "Hello, [name]}".

The `exit 1` command is used to indicate that the script has exited with an error. The number "1" is known as an exit code or exit status.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
