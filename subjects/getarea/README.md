## get-area

### instructions

- Write a program that takes a number and prints the area of the circle.
    - The area of the circle is `π` times the radius squared.
    - if the number is not positive or is not a number, print `Error` followed by (`'\n'`)
    - if the number is positive, print the area of the circle followed by (`'\n'`) 
    - if the number of arguments is not `1` print (`'\n'`)
    - You may assume that the value of `π` is 3.14

###  Usage

```console
$ go run . | cat -e
$
$ go run . 10 | cat -e
315$
$ go run . 4 | cat -e
50$
$ go run . -10 | cat -e
Error$
$ go run . "Hello World!" | cat -e
Error$
```
