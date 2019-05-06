## brackets

### Instructions

Write a program that takes an undefined number of strings in arguments. For each
argument, the program prints on the standard output "OK" followed by a newline
if the expression is correctly bracketed, otherwise it prints "Error" followed by
a newline.

Symbols considered as `brackets` are brackets `(` and `)`, square brackets `[`
and `]`and braces `{` and `}`. Every other symbols are simply ignored.

An opening bracket must always be closed by the good closing bracket in the
correct order. A string which do not contains any bracket is considered as a
correctly bracketed string.

If there is no arguments, the program must print only a newline.

Examples of outputs :

```console
student@ubuntu:~/student/test$ go build
student@ubuntu:~/student/test$ ./test '(johndoe)' | cat -e
OK$
student@ubuntu:~/student/test$ ./test '([)]' | cat -e
Error$
student@ubuntu:~/student/test$ ./test '' '{[(0 + 0)(1 + 1)](3*(-1)){()}}' | cat -e
OK$
OK$
student@ubuntu:~/student/test$ ./test | cat -e
$
student@ubuntu:~/student/test$
```
