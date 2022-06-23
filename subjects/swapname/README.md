## swap-name

### Instructions

Write a program that takes two words in the first argument and swaps them.
- If the number of arguments is not 2, nothing should be printed.
- The first argument should contain only alphabetic characters and spaces and only two word, if not, print `"Error\n"`
- Print the result with a new line at the end (`'\n'`).
- Trim space in the beginning and end of the string

### Usage

```console
$ go run . | cat -e
$ go run . "Hello World" "Hello World" | cat -e
$ go run . "Hello World" | cat -e
World Hello$
$ go run . "He2 $%hello" | cat -e
Error$
$ go run . "    Hello World   " | cat -e
World Hello$
$ go run . "Hello World Hello" | cat -e
Error$
$ go run . "Hello" | cat -e
Error$
```
