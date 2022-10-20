## getalpha

### Instructions

Write  a program that takes an `int` [0-127] as an argument and returns the corresponding ASCII table character, followed by newline.

- If the integer above 127 or less than 0 returns newline.
- If it's not a number returns newline.
- If the number of arguments is above 1 return newline.

### Usage


```go
$ go run . | cat -e
$
$ go run . "98" | cat -e
a
$ go run . "95" | cat -e
;
$ go run . "95"  "102" | cat -e
$
```