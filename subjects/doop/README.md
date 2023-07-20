## doop

### Instructions

Write a program that is called `doop`.

The program has to be used with three arguments:

- A value
- An operator, one of : `+`, `-`, `/`, `*`, `%`
- Another value

In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

The program has to handle the modulo and division operations by 0 as shown on the output examples below.

### Usage

```console
$ go run .
$ go run . 1 + 1 | cat -e
2
$
$ go run . hello + 1
$ go run . 1 p 1
$ go run . 1 / 0 | cat -e
No division by 0
$
$ go run . 1 % 0 | cat -e
No modulo by 0
$
$ go run . 9223372036854775807 + 1
$ go run . -9223372036854775809 - 3
$ go run . 9223372036854775807 "*" 3
$ go run . 1 "*" 1
1
$ go run . 1 "*" -1
-1
$
```

### Notions

- [Numeric Types](https://golang.org/ref/spec#Numeric_types)
- [Arithmetic Operators](https://golang.org/ref/spec#Arithmetic_operators)
