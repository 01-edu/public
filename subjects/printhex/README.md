## printhex

### Instructions

Write a program that takes a positive (or zero) number expressed in base 10, and that displays it in base 16 (with lowercase letters) followed by a newline (`'\n'`).

- If the number of arguments is different from 1, the program displays nothing.
- Error cases have to be handled as shown in the example below.

### Usage

```console
$ go run . 10
a
$ go run . 255
ff
$ go run . 5156454
4eae66
$ go run .
$ go run . "123 132 1" | cat -e
ERROR$
$
```
