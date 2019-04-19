# cleanstr
## Instructions

Write a program that takes a string, and displays this string with exactly one
space between words, with no spaces or tabs either at the beginning or the end,
followed by a \n.

A "word" is defined as a part of a string delimited either by spaces/tabs, or
by the start/end of the string.

If the number of arguments is not 1, or if there are no words to display, the
program displays \n.

And its output :

```console
student@ubuntu:~/student/cleanstr$ go build
student@ubuntu:~/student/cleanstr$ ./cleanstr "you see it's easy to display the same thing" | cat -e
you see it's easy to display the same thing$
student@ubuntu:~/student/cleanstr$ ./cleanstr " only    it's  harder   " 
only it's harder$
student@ubuntu:~/student/cleanstr$ ./cleanstr " how funny" "Did you   hear Mathilde ?"
$
student@ubuntu:~/student/cleanstr$ ./cleanstr "" | cat -e 
$
student@ubuntu:~/student/cleanstr$
```
