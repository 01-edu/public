## brackets

### Instructions

Write a program that takes an undefined number of `strings` in arguments. For each
argument, if the expression is correctly bracketed, the program prints on the standard output `OK` followed by a newline(`'\n'`), otherwise it prints "Error" followed by
a newline.

Symbols considered as `brackets` are parentheses `(` and `)`, square brackets `[`
and `]` and curly braces `{` and `}`. Every other symbols are simply ignored.

An opening bracket must always be closed by the good closing bracket in the
correct order. A `string` which does not contain any bracket is considered as a
correctly bracketed `string`.

If there is no argument, the program must print only a newline.

### Usage

```console
student@ubuntu:~/student/brackets$ go build
student@ubuntu:~/student/brackets$ ./brackets '(johndoe)' | cat -e
OK$
student@ubuntu:~/student/brackets$ ./brackets '([)]' | cat -e
Error$
student@ubuntu:~/student/brackets$ ./brackets '' '{[(0 + 0)(1 + 1)](3*(-1)){()}}' | cat -e
OK$
OK$
student@ubuntu:~/student/brackets$ ./brackets | cat -e
$
student@ubuntu:~/student/brackets$

```
