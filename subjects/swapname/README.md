## swap-name

### Instructions

Write a program that takes two words in the first argument and swaps them.
- If the number of arguments is not 2, nothing should be printed.
- If the first argument is not an alphanumeric characters, the program should print `"Error\n"`.
- Skip the spaces in the first of the two words.
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
