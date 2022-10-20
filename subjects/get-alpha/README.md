# getalpha

### Instructions

write  a program that takes an integer [0-127] in parameter and returns the char that match the index of the integer in the ASCII table follwed by newline, follwed by newline.

- If the integer above 127 or less than 0 returns newline.
- If it's not a number returns newline.
- If tge number of arguments above 1 return newline.

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