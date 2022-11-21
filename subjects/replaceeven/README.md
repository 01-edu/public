## replaceeven

### Instructions

Write a program that takes every character with an even index of each argument and replaces it with the number `2`.

- Your program should always print a newline `\n` at the end.
- Your program should print all the arguments in one line, separated by the space character `' '`.
- `0` is an even number.

### Usage

```console
$ go run . | cat -e
$
$ go run . hello 444444444444444 | cat -e
2e2l2 242424242424242$
$ go run . "Zone 01 is the best" | cat -e
2o2e202 2s2t2e2b2s2$
```
