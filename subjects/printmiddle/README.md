## print-middle

### Instructions

Write a program that prints the middle argument of the command line. 
- If the number of arguments is less than 1, print (`'\n'`)
- If the number of arguments is **even**, print the middle two arguments with a space (`' '`) between the arguments
- If the number of arguments is **odd**, print the middle one.
- Print (`'\n'`) at the end of the output.

### Usage

```console
$ go run . | cat -e
$
$ go run . "Hello World!" | cat -e
Hello World!$
$ go run . "Hello World!" "World!" "All" | cat -e
World!$
$ go run . 1 2 3 4 | cat -e
2 3$
$ go run . "Hello" "World!" | cat -e 
Hello World!$
```
