## gcd

### Instructions

Write a program that takes two `string` representing two strictly positive integers that fit in an `int`.

The program displays their greatest common divisor followed by a newline (`'\n'`).

If the number of arguments is different from 2, the program displays nothing.

All arguments tested will be positive `int` values.

### Usage

```console
student@ubuntu:~/gcd$ go build
student@ubuntu:~/gcd$ ./gcd 42 10 | cat -e
2$
student@ubuntu:~/gcd$ ./gcd 42 12
6
student@ubuntu:~/gcd$ ./gcd 14 77
7
student@ubuntu:~/gcd$ ./gcd 17 3
1
student@ubuntu:~/gcd$ ./gcd
student@ubuntu:~/gcd$ ./gcd 50 12 4
student@ubuntu:~/gcd$
```
