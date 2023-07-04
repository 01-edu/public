# getascii

### Instructions

Write a program that takes a character as an argument, and displays the correspondent ASCII code followed by a newline.

- If the input is more than one character, return a newline `\n`.
- If the number of arguments is less than 1, return a newline `\n`.
- If the number of arguments is more than 1, return a newline `\n`.
- If the the input is a non ascii character, return a newline `\n`.

### Usage

```console
console
$ go run . "W" | cat -e
87$
$
$ go run . "W" "Q" "T" | cat -e
$
$
$ go run . "^" | cat -e
94$
$
$ go run . 7 | cat -e
55$
$
$ go run . a | cat -e
97$
$
$ go run . "Stay Home" | cat -e
$
$ go run .
$
```
