## Same-Case 

### Instructions

Write a program that takes 2 characters as arguments and does the following:

- If the number of arguments are more than two arguments return newline.
- If either of the characters is not a letter, return -1 followed by a newline.
- If both characters are the same case, return 1 followed by a newline.
- If both characters are letters but not the same case, return 0 followed by a newline.
- If the argument is more than one character or if it's empty, return a newline.

### Usage

```console
$ go run . | cat -e
$
$ go run . "p" "w"| cat -e
1$
$ go run . "p" "Q" | cat -e
0$
$ go run . "p" "4"  | cat -e
-1$
$ go run . "P" "Q" "W" | cat -e
$
$ go run . "A" "b"  | cat -e
0$
```
