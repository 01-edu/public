## arg-issorted

### Instruction

Write a program where you define if an argument is sorted or not sorted:
- Your Program should return T if an argument is sorted.
- Your Program should return F if an argument is not sorted.
- If The argument is empty nothing will be printed.
- If the number of arguments is not `1` Print nothing.
- If the length argument is equal to `1` Print nothing.
- Your program should always print `('\n') `at the end of the output.

### Usage

Here is a program to test your function.

```console
$ go run . | cat -e
$
$ go run . a | cat -e
$
$ go run . " 5ABc" | cat -e
T$
$ go run . "zA1" | cat -e
F$
$ go run . "01Talent" "student" | cat -e
$
$ go run . "Hello World" | cat -e
F$
$ go run .  "a b c" | cat -e
F$
```

