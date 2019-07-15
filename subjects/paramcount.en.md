## paramcount

### Instructions

Write a program that displays the number of arguments passed to it. This number will be followed by
a newline(`'\n'`).

If there is no argument, the program displays `0` followed by a newline.

### Usage

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
