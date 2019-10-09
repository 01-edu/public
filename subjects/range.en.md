## range

### Instructions

Write a program which must:

-   **Allocate (with `make`)** an array of integers.

-   Fill it with consecutive values that begins at the first argument and end at the second argument (Including the values of thoses arguments !).

-   That prints the array.

Errors should be handled.

If the number of arguments is different from 2 the program prints a newline ("`\n`").

### Expected output :

```console
student@ubuntu:~/range$ go build
student@ubuntu:~/range$ ./range 1 3
[1 2 3]
student@ubuntu:~/range$ ./range -1 2 | cat -e
[-1 0 1 2]$
student@ubuntu:~/range$ ./range 0 0
[0]
student@ubuntu:~/reverserange$ ./reverserange 0 nan | cat -e
strconv.Atoi: parsing "nan": invalid syntax$
student@ubuntu:~/range$
```
