## reverserange

### Instructions

Write a program that takes two numbers as arguments and prints the consecutive values that begins at the second number and end at the first number (including the values of those numbers !).

Errors should be handled.

If the number of arguments is different from 2 the program prints nothing.

### Usage :

```console
student@ubuntu:~/reverserange$ go build
student@ubuntu:~/reverserange$ ./reverserange 1 3
3 2 1
student@ubuntu:~/reverserange$ ./reverserange -1 2 | cat -e
2 1 0 -1$
student@ubuntu:~/reverserange$ ./reverserange 0 0
0
student@ubuntu:~/reverserange$ ./reverserange 0 -3
-3 -2 -1 0
student@ubuntu:~/reverserange$ ./reverserange 0 nan | cat -e
strconv.Atoi: parsing "nan": invalid syntax$
student@ubuntu:~/reverserange$
```
