## getalpha

### Instructions

Write a program that takes an `int` between `[0-127]` as an argument and returns the corresponding ASCII table character, followed by a newline `\n`.

- If the integer above 127 or less than 0 return a newline `\n`.
- If it's not a number return a newline `\n`.
- If the number of arguments is above 1 return a newline `\n`.

### Usage

```go
$ go run . | cat -e
$
$ go run . "98" | cat -e
b
$ go run . "95" | cat -e
_
$ go run . "95"  "102" | cat -e
$
```
