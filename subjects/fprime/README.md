## fprime

### Instructions

Write a program that takes a positive `int` and displays its prime factors, followed by a newline (`'\n'`).

- Factors must be displayed in ascending order and separated by `*`.

- If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

### Usage

```console
$ go run . 225225
3*3*5*5*7*11*13
$ go run . 8333325
3*3*5*5*7*11*13*37
$ go run . 9539
9539
$ go run . 804577
804577
$ go run . 42
2*3*7
$ go run . a
$ go run . 0
$ go run . 1
$
```
