## range

### Instructions

Write a program that takes two numbers as arguments and prints the consecutive values that begins at the first number and end at the second number (including the values of those numbers !).

Errors should be handled.

If the number of arguments is different from 2 the program prints nothing.

### Usage :

```console
student@ubuntu:~/range$ go build
student@ubuntu:~/range$ ./range 1 3
1 2 3
student@ubuntu:~/range$ ./range -1 2 | cat -e
-1 0 1 2$
student@ubuntu:~/range$ ./range 0 0
0
student@ubuntu:~/range$ ./range 0 nan | cat -e
strconv.Atoi: parsing "nan": invalid syntax$
student@ubuntu:~/range$
```
