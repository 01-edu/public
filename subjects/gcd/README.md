## gcd

### Instructions

Write a program that takes two `string` representing two strictly positive integers that fit in an `int`.

The program displays their greatest common divisor followed by a newline (`'\n'`).

If the number of arguments is different from 2, the program displays nothing.

All arguments tested will be positive `int` values.

### Usage

```console
$ go run . 42 10 | cat -e
2$
$ go run . 42 12
6
$ go run . 14 77
7
$ go run . 17 3
1
$ go run .
$ go run . 50 12 4
$
```
