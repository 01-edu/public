## printchessboard

### Instructions

Write a program that takes two integers as arguments and displays the chess desk, in which white cells are represented by `' '` and black cells by `'#'`.

- If the number of arguments is different from 2, or if the argument is not a positive number, the program displays `Error` followed by a newline (`'\n'`).

### Usage

```console
student@ubuntu:~/printchessboard$ go build
student@ubuntu:~/printchessboard$ ./printchessboard 4 3 | cat -e
# # $
 # #$
# # $
student@ubuntu:~/printchessboard$ ./printchessboard 7 | cat -e
Error$
student@ubuntu:~/printchessboard$ ./printchessboard 0 0 | cat -e
Error$
student@ubuntu:~/printchessboard$
```
