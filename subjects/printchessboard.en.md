## printchessboard

### Instructions

Write a program that takes a two integer as argument and displays the chess desk, where white cells `' '` and black cells `'#'`.

-   If the number of arguments is different from 2, or if the argument is not a positive number, the program displays `Error` followed by a newline (`'\n'`).

### Usage

```console
student@ubuntu:~/[[ROOT]]/printchessboard$ go build
student@ubuntu:~/[[ROOT]]/printchessboard$ ./printchessboard 4 3 | cat -e
# # $
 # #$
# # $
student@ubuntu:~/[[ROOT]]/printchessboard$ ./printchessboard 7 | cat -e
Error$
student@ubuntu:~/[[ROOT]]/printchessboard$
```
