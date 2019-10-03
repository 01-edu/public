## printmemory

### Instructions

Write a program that takes as arguments a set of 10 numbers and displays the memory as in the example.

- If the number of arguments 0 or bigger than 10 in should print a newline `\n`.

- In case one of the arguments is not a number you should handle the error.

### Expected output :

```console
student@ubuntu:~/foreachprog$ go build
student@ubuntu:~/foreachprog$ ./foreachprog 104 101 108 108 111 16 21 42 | cat -e
6800 0000 6500 0000 6c00 0000 6c00 0000 $
6f00 0000 1000 0000 1500 0000 2a00 0000 $
0000 0000 0000 0000 $
hello..*..$
student@ubuntu:~/foreachprog$
```
