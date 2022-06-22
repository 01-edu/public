## swap-name

### Instructions

Write a program that takes two words in the first argument and swaps them.
- If the number of arguments is not 2, nothing should be printed.
- The first argument should contain only alphabetic characters and spaces, if not, print `"Error\n"`
- Skip the spaces only in the first word of the first argument.
- Print the result with a new line at the end (`'\n'`).

### Usage

```console
$ go run . | cat -e
$ go run . "Hello World" "Hello World" | cat -e
$ go run . "Hello World" | cat -e
World Hello$
$ go run . "He2 $%hello" | cat -e
Error$
$ go run . "   Hello World" | cat -e
World Hello$
$ go run . "Hello World   " | cat -e
Error$
$ go run . "Hello World Hello" | cat -e
Error$
```
