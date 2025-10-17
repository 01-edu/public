## Calculator

### Instructions

In this exercise you will make a script `calculator.sh` that will take 3 arguments, calculate the result and print it on the standard output.

- The first argument and the third argument will be numbers.
- The second argument will be the operator.

Each operator should have its own function named as follow:

- `+`: `do_add()`.
- `-`: `do_sub()`.
- `*`: `do_mult()`.
- `/`: `do_divide()`.

Each function will receive two arguments, the left number and the right number, and return the result of the operation.

The functions assume that the input is valid, so the input must be checked before calling the functions.

To choose which function to call you must use the `case` statement.

> The functions will also be tested individually, so it is important to name each function exactly as above, the behavior of the functions have to match the exercise instructions.

### Usage

```console
$ ./calculator.sh 20 "*" 3
60
$ ./calculator.sh 20 / 20
1
$ ./calculator.sh -1 - 10
-11
$
```

### Error handling

All errors will print a specific message on **stderr** (ending with a newline) and returns a specific non-zero value:

- Wrong number of arguments: `"Error: expect 3 arguments"`, returns `1`.
- Division by 0: `"Error: division by 0"`, exit with `2`.
- Invalid operator: `"Error: invalid operator"`, exit with `3`.
- Invalid number(s): `"Error: invalid number"`, exit with `4`.

> Negative numbers are also a valid input.

### Hints

- `case` statement example:

```sh
# Check the first argument given to a script
case $1 in
    "left")
        echo "We will turn left"
        ;;

    "right")
        echo "We will turn right"
        ;;

    "top")
        echo "We will turn top"
        ;;

    "bottom")
        echo "We will turn bottom"
        ;;

    # Any other case
    *)
        # This is printed in stderr
        >&2 echo "Error: invalid argument"
        exit 2
        ;;
esac
```

- Example of a function taking two arguments and returning a value by printing it.
  The behavior of this function is the same than the one expected for the operators functions you will create:

```sh
print_full_name () {
    name=$1
    surname=$2
    echo $name $surname
}

print_full_name "Gene" "Mallamar"
```

> Google and Man will be your friends!

### References

- [Bash functions](https://linuxize.com/post/bash-functions/)
- [Test if a variable is a number](https://stackoverflow.com/questions/806906/how-do-i-test-if-a-variable-is-a-number-in-bash)
- [Print on standard error](https://stackoverflow.com/questions/2990414/echo-that-outputs-to-stderr)
- [Case statement](https://linuxize.com/post/bash-case-statement/)
