## gcd

### Instructions

Write a program that takes two `string` representing two strictly positive integers that fit in an `int`.

The program displays their greatest common divisor followed by a newline (`'\n'`).

If the number of parameters is different from 2, the program displays nothing.

All arguments tested will be positive `int` values.

### Usage

```console
student@ubuntu:~/[[ROOT]]/gcd$ go build
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 42 10 | cat -e
2$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 42 12
6
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 14 77
7
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 17 3
1
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 50 12 4
student@ubuntu:~/[[ROOT]]/gcd$
```
