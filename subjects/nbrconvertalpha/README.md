## nbrconvertalpha

### Instructions

Write a **program** that prints the corresponding letter in the `n` position of the latin alphabet, where `n` is each argument received.

For example `1` matches `a`, `2` matches `b`, etc. If `n` does not match a valid position of the alphabet or if the argument is not an integer, the **program** should print a space (" ").

A flag `--upper` should be implemented. When used the program prints the result in upper case. The flag will always be the first argument.

### Usage

```console
student@ubuntu:~/student/test$ go build
student@ubuntu:~/student/test$ ./nbrconvertalpha
student@ubuntu:~/student/test$ ./nbrconvertalpha 8 5 12 12 15 | cat -e
hello$
student@ubuntu:~/student/test$ ./nbrconvertalpha 12 5 7 5 14 56 4 1 18 25 | cat -e
legen dary$
student@ubuntu:~/student/test$ ./nbrconvertalpha 32 86 h | cat -e
   $
student@ubuntu:~/student/test$ ./nbrconvertalpha --upper 8 5 25
HEY$
student@ubuntu:~/student/test$
```
