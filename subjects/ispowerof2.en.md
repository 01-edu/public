## ispowerof2

### Instructions

Write a program that determines if a given number is a power of 2.

This program must print `true` if the given number is a power of 2, otherwise it prints `false`.

-   If there is more than one or no argument the program should print a newline ("`\n`").

-   Error cases have to be handled.

### Usage :

```console
student@ubuntu:~/ispowerof2$ go build
student@ubuntu:~/ispowerof2$ ./ispowerof2 2 | cat -e
true$
student@ubuntu:~/ispowerof2$ ./ispowerof2 64 | cat -e
true$
student@ubuntu:~/ispowerof2$ ./ispowerof2 513 | cat -e
false$
student@ubuntu:~/ispowerof2$ ./ispowerof2

student@ubuntu:~/ispowerof2$ ./ispowerof2 64 1024

student@ubuntu:~/ispowerof2$ ./ispowerof2 notanumber | cat -e
strconv.Atoi: parsing "a": invalid syntax$
student@ubuntu:~/ispowerof2$
```
