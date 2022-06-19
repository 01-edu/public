## print-ascii

### Instructions

Write a program that prints the ASCII value of a letter passed in the command line
- If the argument is not a letter nothing will be printed
- if the number of arguments is not 1 then nothing will be printed

### Usage

```console
$ go run .
$ go run . a
97
$ go run . 'A'
65
$ go run . 'z'
122
$ go run . Z
90
$ go run . 1
$ go run . "Hello" "Word"
```
