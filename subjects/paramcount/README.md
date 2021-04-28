## paramcount

### Instructions

Write a program that displays the number of arguments passed to it. This number will be followed by a newline (`'\n'`).

If there is no argument, the program displays `0` followed by a newline.

### Usage

```console
$ go run . 1 2 3 5 7 24
6
$ go run . 6 12 24 | cat -e
3$
$ go run . | cat -e
0$
$
```
