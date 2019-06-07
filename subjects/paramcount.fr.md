## paramcount

### Instructions

Write a program that displays the number of arguments passed to it. This number will be followed by
a newline.

If there are no arguments, just display a 0 followed by a newline.

Examples of outputs :

```console
student@ubuntu:~/student/paramcount$ go build
student@ubuntu:~/student/paramcount$ ./paramcount 1 2 3 5 7 24
6
student@ubuntu:~/student/paramcount$ ./paramcount 6 12 24 | cat -e
3$
student@ubuntu:~/student/paramcount$ ./paramcount | cat -e
0$
student@ubuntu:~/student/paramcount$
```
