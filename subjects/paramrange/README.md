## param-range

### instructions

Write a program that takes a number in the arguments and prints the max and min.
- If the number of arguments is less than 2 print (`'\n'`)
- If one of the arguments is not a number, print (`"Error\n"`)
- The output should be space-separated and (`'\n'`) at the end.

### Usage
```console
$ go run . | cat -e
$
$ go run . 1 2 3 4 5 6 7 8 9 | cat -e
1 9$
$ go run . "-1" "1" | cat -e
-1 1$
$ go run . 1 a 2 3 4 5 6 7 8 9 | cat -e
Error$
```
