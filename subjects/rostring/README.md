## rostring

### Instructions

Write a program that takes a `string` and displays this `string` after rotating it
one word to the left.

Thus, the first word becomes the last, and others stay in the same order.

A word is a sequence of **alphanumerical** characters.

Words will be separated by only one space in the output.

If the number of arguments is different from 1, the program displays a newline.

### Usage

```console
$ go run . "abc   " | cat -e
abc$
$ go run . "Let there     be light"
there be light Let
$ go run . "     AkjhZ zLKIJz , 23y"
zLKIJz , 23y AkjhZ
$ go run . | cat -e
$
$
```
