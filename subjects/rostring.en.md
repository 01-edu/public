## rostring

### Instructions

Write a program that takes a string and displays this string after rotating it
one word to the left.

Thus, the first word becomes the last, and others stay in the same order.

A "word" is defined as a part of a string delimited either by spaces/tabs, or
by the start/end of the string.

Words will be separated by only one space in the output.

If the number of arguments is not one, the program displays \n.

And its output :

```console
student@ubuntu:~/piscine/rostring$ go build
student@ubuntu:~/piscine/rostring$ ./rostring "abc   " | cat -e
abc$
student@ubuntu:~/piscine/rostring$ ./rostring "Que la     lumiere soit et la lumiere fut"
la lumiere soit et la lumiere fut Que
student@ubuntu:~/piscine/rostring$ ./rostring "     AkjhZ zLKIJz , 23y"
zLKIJz , 23y AkjhZ
student@ubuntu:~/piscine/rostring$ ./rostring | cat -e
$
student@ubuntu:~/piscine/rostring$
```
