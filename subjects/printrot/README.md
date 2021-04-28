## printrot

### Instructions

Write a program that takes a lower case letter and goes around in order to display the whole alphabet starting from
this letter on a single line.

A line is a sequence of characters preceding the [end of line](https://en.wikipedia.org/wiki/Newline) character (`'\n'`).

If the input is invalid the program prints a newline (`'\n'`).

### Usage

```console
$ go run . "abc"

$ go run . "a"
abcdefghijklmnopqrstuvwxyz
$ go run . "g"
ghijklmnopqrstuvwxyzabcdef
$ go run . "a" "a"

$ go run . "A"

$
```
