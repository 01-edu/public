## expandstr

### Instructions

Write a program that takes a string and displays it with exactly three spaces
between each word, with no spaces or tabs either at the beginning or the end,
followed by a newline.

A word is a section of string delimited either by spaces/tabs, or by the
start/end of the string.

If the number of parameters is not 1, or if there are no words, simply display
a newline.

Example of output :

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
