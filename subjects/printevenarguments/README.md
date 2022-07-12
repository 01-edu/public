## printevenarguments

### Instructions

Write a program that receives some arguments from the command line and prints the ones with an even index followed by a new line. If the number of even arguments is 0, print a new line.

### Usage

```console
$ go run . first second third | cat -e
second$
$ go run . a b c d | cat -e
b$
d$
$ go run . a | cat -e
$
```
