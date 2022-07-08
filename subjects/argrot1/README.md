## arg-rot1

### Instructions

Write a program that rotates the arguments in the command line by 1 position.
- If the number of arguments is less than 2, print (`'\n'`).
- Print (`'\n'`) at the end of the output.

### Usage

```console
$ go run . | cat -e
$
$ go run . "hello" | cat -e
$
$ go run . 1 2 3 4 5 6 7 8 9 | cat -e
2 3 4 5 6 7 8 9 1$
$ go run "Hello" "World" | cat -e
World Hello$
$ go run . a b c d
b c d a$
```
