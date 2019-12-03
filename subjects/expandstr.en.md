## expandstr

### Instructions

Write a program that takes a `string` and displays it with exactly three spaces
between each word, with no spaces or tabs at either the beginning nor the end.

The `string` will be followed by a newline (`'\n'`).

A word is a sequence of alphanumerical characters.

If the number of parameters is not 1, or if there are no word, the program displays
a newline.

### Usage

```console
student@ubuntu:~/[[ROOT]]/expandstr$ go build
student@ubuntu:~/[[ROOT]]/expandstr$ ./expandstr "you   see   it's   easy   to   display   the   same   thing" | cat -e
you   see   it's   easy   to   display   the   same   thing$
student@ubuntu:~/[[ROOT]]/expandstr$ ./expandstr "   only  it's harder   " | cat -e
only   it's   harder$
student@ubuntu:~/[[ROOT]]/expandstr$ ./expandstr " how funny it is" "did you  hear, Mathilde ?" | cat -e
$
student@ubuntu:~/[[ROOT]]/expandstr$ ./expandstr | cat -e
$
student@ubuntu:~/[[ROOT]]/expandstr$
```
