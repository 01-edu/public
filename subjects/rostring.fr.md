## rostring

### Instructions

Write a program that takes a string and displays this string after rotating it
one word to the left.

Thus, the first word becomes the last, and others stay in the same order.

A word is a sequence of **alphanumerical** characters.

Words will be separated by only one space in the output.

If the number of arguments is not one, the program displays a newline.

Examples of outputs :

```console
student@ubuntu:~/piscine/rostring$ go build
student@ubuntu:~/piscine/rostring$ ./rostring "abc   " | cat -e
abc$
student@ubuntu:~/piscine/rostring$ ./rostring "Let there     be light"
there be light There
student@ubuntu:~/piscine/rostring$ ./rostring "     AkjhZ zLKIJz , 23y"
zLKIJz , 23y AkjhZ
student@ubuntu:~/piscine/rostring$ ./rostring | cat -e
$
student@ubuntu:~/piscine/rostring$
```
