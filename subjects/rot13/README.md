## rot13

### Instructions

Write a program that takes a `string` and displays it, replacing each of its
letters by the letter 13 spots ahead in alphabetical order.

- 'z' becomes 'm' and 'Z' becomes 'M'. The case of the letter stays the same.

- The output will be followed by a newline (`'\n'`).

- If the number of arguments is different from 1, the program displays nothing.

### Usage

```console
$ go run . "abc"
nop
$ go run . "hello there"
uryyb gurer
$ go run . "HELLO, HELP"
URYYB, URYC
$ go run .
$
```
