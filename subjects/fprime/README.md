## fprime

### Instructions

Write a program that takes a positive `int` and displays its prime factors, followed by a newline (`'\n'`).

- Factors must be displayed in ascending order and separated by `*`.

- If the number of arguments is different from 1, if the argument is invalid or if the integer doesn't have a prime factor, the program displays nothing.

### Usage

```console
student@ubuntu:~/fprime/test$ go build
student@ubuntu:~/fprime/test$ ./test 225225
3*3*5*5*7*11*13
student@ubuntu:~/fprime/test$ ./test 8333325
3*3*5*5*7*11*13*37
student@ubuntu:~/fprime/test$ ./test 9539
9539
student@ubuntu:~/fprime/test$ ./test 804577
804577
student@ubuntu:~/fprime/test$ ./test 42
2*3*7
student@ubuntu:~/fprime/test$ ./test a
student@ubuntu:~/fprime/test$ ./test 0
student@ubuntu:~/fprime/test$ ./test 1
student@ubuntu:~/fprime/test$
```
