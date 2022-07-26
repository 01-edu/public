## swap-name

### Instructions

Write a program that takes a first and last name as one argument and swaps them.
- If there is more than one argument, nothing should be printed.
- The argument should contain only alphabetic characters and spaces and only two words, if not, you have to print `"Error\n"`
- The result must be followed with a new line at the end (`'\n'`).
- Trim any extra spaces in the string.

### Usage

```console
$ go run . | cat -e
$ go run . "James Carl" "Peter Lab" | cat -e
$ go run . "ROB carlos" | cat -e
carlos ROB$
$ go run . "He2 $%hello" | cat -e
Error$
$ go run . "    JULIE BIRD   " | cat -e
BIRD JULIE$
$ go run . "Carl Jon Mic" | cat -e
Error$
$ go run . "ROB" | cat -e
Error$
$ go run . "ROB       FASEM" | cat -e
FASEM ROB$
```
