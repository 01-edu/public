## division

### Instructions

In this exercise, you will make a script `division.sh` that will take two arguments from the command line, and divide the first one by the second one.

If there is a remainder after doing the division, it should be ignored.

You will need to handle what to do when the inputs are wrong:

- If the divisor is `0` you will need to output `Error: division by zero is not allowed`.

- If the arguments are not numbers, the output should be `Error: both arguments must be numeric`.

- In the case where the number of arguments are not enough, the output should be `Error: two numbers should be provided`.

Your script should handle very large numbers as well.

For this exercise the use of the `test` command is not allowed.

### Usage

```console
$ ./divide.sh 4 1
4
$
```

### Hints

You can use the following to help you solve this exercise:

[Bash conditional construct](https://www.gnu.org/software/bash/manual/bash.html#Conditional-Constructs) can be used to decide whether to execute a specific command. Below an example `script.sh`.

```bash
#!/usr/bin/env bash
if [[ 1 > 2 ]]; then
    echo "true"
else
    echo "false"
fi
```

And its output:

```console
$ bash script.sh
false
```

It is possible to combine several conditions with the **AND** (`&&`) and **OR** (`||`) logical operators. Below and example `script.sh`.

```bash
if [[ 1 > 2 ]] || [[ 1 == 1 ]]; then
    echo "true"
else
    echo "false"
fi
if [[ 1 > 2 ]] && [[ 1 == 1 ]]; then
    echo "true"
else
    echo "false"
fi
```

And its output:

```console
true
false
```

[bc](https://www.gnu.org/software/bc/manual/html_mono/bc.html) is a Unix utility that performs arbitrary precision arithmetic. It is particularly useful to handle numbers that are too large. One way of using it is as below:

```console
$ echo "2 + 2" | bc
4
$
```
