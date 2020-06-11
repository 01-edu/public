## printhex

### Instructions

Write a program that takes a positive (or zero) number expressed in base 10, and displays it in base 16 (with lowercase letters) followed by a newline (`'\n'`).

- If the number of arguments is different from 1, the program displays nothing.
- Error cases have to be handled as shown in the example below.

### Usage

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test 10
a
student@ubuntu:~/[[ROOT]]/test$ ./test 255
ff
student@ubuntu:~/[[ROOT]]/test$ ./test 5156454
4eae66
student@ubuntu:~/[[ROOT]]/test$ ./test
student@ubuntu:~/[[ROOT]]/test$ ./test "123 132 1" | cat -e
ERROR$
student@ubuntu:~/[[ROOT]]/test$
```
