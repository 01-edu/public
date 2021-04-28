## printchessboard

### Instructions

Write a program that takes two integers as arguments and displays the chess desk, in which white cells are represented by `' '` and black cells by `'#'`.

- If the number of arguments is different from 2, or if the argument is not a positive number, the program displays `Error` followed by a newline (`'\n'`).

### Usage

```console
$ go run . 4 3 | cat -e
# # $
 # #$
# # $
$ go run . 7 | cat -e
Error$
$ go run . 0 0 | cat -e
Error$
$
```
