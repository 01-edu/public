## get-area

### Instructions

Write a program that takes a positive number as radius and prints the area of a circle.
  - The area of the circle is `3.14` times the radius squared.
  - Only positive numbers will be accepted, otherwise print `Error` followed by (`'\n'`)
  - If the number of arguments is not `1` print (`'\n'`)
  - The output must be an integer.


###  Usage

```console
$ go run . | cat -e
$
$ go run . 10 | cat -e
314$
$ go run . 4 | cat -e
50$
$ go run . -10 | cat -e
Error$
$ go run . "Hello World!" | cat -e
Error$
```
