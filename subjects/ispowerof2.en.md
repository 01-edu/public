## ispowerof2

### Instructions

Write a program that determines if a given number is a power of 2.

This program must print `true` if the given number is a power of 2, otherwise it prints `false`.

- If there's no arguments passed to the program or there's more that one it shohuld print `\n`.

- In case of error you shoud handle it.

### Expected output :

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

student@ubuntu:~/ispowerof2$
```