## fprime

### Instructions

Write a program that takes a positive `int` and displays its prime factors, followed by a newline (`'\n'`).

- Factors must be displayed in ascending order and separated by `*`.

- If the number of arguments is different from 1, if the argument is invalid or if the integer doesn't have a prime factor, the program displays nothing.

### Usage

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test 225225
3*3*5*5*7*11*13
student@ubuntu:~/[[ROOT]]/test$ ./test 8333325
3*3*5*5*7*11*13*37
student@ubuntu:~/[[ROOT]]/test$ ./test 9539
9539
student@ubuntu:~/[[ROOT]]/test$ ./test 804577
804577
student@ubuntu:~/[[ROOT]]/test$ ./test 42
2*3*7
student@ubuntu:~/[[ROOT]]/test$ ./test a
student@ubuntu:~/[[ROOT]]/test$ ./test 0
student@ubuntu:~/[[ROOT]]/test$ ./test 1
student@ubuntu:~/[[ROOT]]/test$
```
