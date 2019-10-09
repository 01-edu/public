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
student@ubuntu:~/piscine-go/expandstr$ go build
student@ubuntu:~/piscine-go/expandstr$ ./expandstr "you   see   it's   easy   to   display   the   same   thing" | cat -e
you   see   it's   easy   to   display   the   same   thing$
student@ubuntu:~/piscine-go/expandstr$ ./expandstr "   only  it's harder   " | cat -e
only   it's   harder$
student@ubuntu:~/piscine-go/expandstr$ ./expandstr " how funny it is" "did you  hear, Mathilde ?" | cat -e
$
student@ubuntu:~/piscine-go/expandstr$ ./expandstr | cat -e
$
student@ubuntu:~/piscine-go/expandstr$
```
