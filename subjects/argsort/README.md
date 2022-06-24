## argsort

### Instructions

Write a program that checks if an argument is sorted or not.
- Your Program should return `T` followed by a newline (`'`\n'`) if an argument is sorted.
- Your Program should return `F` followed by a newline (`'`\n'`) if an argument is not sorted.


### Usage

And it's output should be:

```console
$ go run . 
$ go run . a | cat -e
$ go run . " 5ABc" | cat -e
T$
$ go run . "zA1" | cat -e
F$
$ go run . "01Talent" "student" | cat -e
$ go run . "Hello World" | cat -e
F$
$ go run .  "a b c" | cat -e
F$
```

