### rev-arg

### Instructions

Write a program that takes the command line arguments and prints them in reverse order.
- if the number of arguments is less than 1, print (`'\n'`)

### Usage

```console
$ go run . | cat -e
$
$ go run . "Hello World!" | cat -e
Hello World!$
$ go run . "Hello World!" "World!" "All" | cat -e
All World! Hello World!$
$ go run . 1 2 3 4 | cat -e
4 3 2 1$
```
