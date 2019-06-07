## expandstr

### Instructions

Write a program that takes a string and displays it with exactly three spaces
between each word, with no spaces or tabs at either the beginning nor the end.

The string will be followed by a newline.

A word is a sequence of alphanumerical characters.

If the number of parameters is not 1, or if there are no words, the program displays
a newline.

Examples of outputs :

```console
student@ubuntu:~/student/expandstr$ go build
student@ubuntu:~/student/expandstr$ ./expandstr "you   see   it's   easy   to   display   the   same   thing" | cat -e
you   see   it's   easy   to   display   the   same   thing$
student@ubuntu:~/student/expandstr$ ./expandstr "   only  it's harder   " | cat -e
only   it's   harder$
student@ubuntu:~/student/expandstr$ ./expandstr " how funny it is" "did you  hear, Mathilde ?" | cat -e
$
student@ubuntu:~/student/expandstr$ ./expandstr | cat -e
$
student@ubuntu:~/student/expandstr$
```
