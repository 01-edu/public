## sumthemall

### Instructions

Create a program that takes `int`´s as arguments and returns the sum of all of them. If it is not a number, there are not enough arguments, or if there is an overflow, print `0` followed by a newline `\n`.

### Usage

```console
$ go run . 23 34 53 | cat -e
110$
$ go run . 1000 2000 3000 4000 | cat -e
10000$
$ go run . a | cat -e
0$
```
