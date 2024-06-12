## concat2args

### Instructions

Write a program that concatenates two arguments and prints the result:

- If the number of arguments is not 2, return a newline.
- If only one of the arguments is empty, print the other.
- If the two arguments are empty return a newline.  

### Usage

```console
$ go run . | cat -e
$
$ go run . "Hello" "World!" | cat -e
HelloWorld!$
$ go run . "01" "talent" | cat -e
01talent$
$ go run . " student " "talented" | cat -e
 student talented$
$ go run . "Hello" "student" "talented" "!" | cat -e
$
``` 
