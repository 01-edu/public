## reverserange

### Instructions

Write a program which must:

-   **Allocate (with `make`)** an array of integers.

-   Fill it with consecutive values that begins at the second argument and end at the first argument (Including the values of thoses arguments !).

-   That prints the array.

Errors should be handled.

If the number of arguments is different from 2 the program prints a newline ("`\n`").

### Usage :

```console
student@ubuntu:~/reverserange$ go build
student@ubuntu:~/reverserange$ ./reverserange 1 3
[3 2 1]
student@ubuntu:~/reverserange$ ./reverserange -1 2 | cat -e
[2 1 0 -1]$
student@ubuntu:~/reverserange$ ./reverserange 0 0
[0]
student@ubuntu:~/reverserange$ ./reverserange 0 -3
[-3 -2 -1 0]
student@ubuntu:~/reverserange$ ./reverserange 0 nan | cat -e
strconv.Atoi: parsing "nan": invalid syntax$
student@ubuntu:~/reverserange$
```
