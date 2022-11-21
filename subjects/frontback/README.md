## frontback

### Instructions

Write a program that prints the previous letter and the next letter of the Latin alphabet.

- If the number of arguments is different from `1`, print a new line `\n`.
- If the length of the argument is different from `1`, print a new line `\n`.
- If the argument is not a letter, print a new line `\n`.
- There are no letters to print before `a` or after `z`.

### Usage

```console
$ go run . | cat -e
$
$ go run . a | cat -e
ab$
$ go run . "B" | cat -e
ABC$
$ go run . "z" | cat -e
yz$
$ go run . "Hello World!" | cat -e
$
$ go run . "1" | cat -e
$
```
