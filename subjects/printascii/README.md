## print-ascii

### instructions

Write a program that prints the ASCII value of a character passed as an argument to the program
- If the argument is not a letter nothing will be printed
- if the number of arguments is not 1 then nothing will be printed
- the program should print (`'\n'`) at the end of the output

### Usage

```console
$ go run . | cat -e
$
$ go run . a | cat -e
97$
$ go run . 'A' | cat -e
65$
$ go run . 'z' | cat -e
122$
$ go run . Z | cat -e
90$
$ go run . 1 | cat -e
$
$ go run . "Hello" "Word" | cat -e
$
```
