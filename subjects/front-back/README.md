## frontback

### Instructions

Write a program that prints the previous and next alphabet of the alphabet.
- If the number of arguments is different from 1, print (`'\n'`).
- If the length of the argument is different from 1, print (`'\n'`)
- If the argument is not a letter, print (`'\n'`)
- There are no letters to print before 'a' or after 'z'.

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
