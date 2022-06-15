## Same-Case 

### Instructions

Write a program that takes 2 arguments and returns:

- If either of the characters is not a letter, return -1
- If both characters are the same case, return 1
- If both characters are letters, but not the same case, return 0
- All outputs should be followed by a new-line character
  
### Usage

```console
$ go run . | cat -e
$
$ go run . "p" "w"| cat -e
1$
$ go run . "p" "Q" | cat -e
0$
$go run . "p" "4"  | cat -e
-1$
$ go run . "P" "Q" "W" | cat -e
$
$go run . A b  | cat -e
-1$
```