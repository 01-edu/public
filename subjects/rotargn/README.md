## rot-arg-n

### Instructions

Write a program that takes arguments from the command line and rotates them by the number in the first argument.
- If the number of arguments is less than 3 print (`'\n'`)
- If the first argument is not a number, the program should print `"Error\n"`
- If the first argument is negative, the program should print `"Error\n"`
- If the first argument is zero, the program should print the original arguments
- Prints the rotated arguments with a space between each argument and a newline at the end.

### Usage

```console
$ go run . | cat -e
$
$ go run . 1 2 3 4 5 | cat -e
3 4 5 2$
$ go run . 2 "Hello" "World" | cat -e
Hello World$
$ go run . -1 2 3 4 5 | cat -e
Error$
$ go run . 0 2 3 4 5 | cat -e
2 3 4 5$
$ go run . 3 "Hello" "World" | cat -e
World Hello$
```