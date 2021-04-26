## rot13

### Instructions

Write a program that takes a `string` and displays it, replacing each of its
letters by the letter 13 spaces ahead in alphabetical order.

- 'z' becomes 'm' and 'Z' becomes 'M'. The case of the letter stays the same.

- The output will be followed by a newline (`'\n'`).

- If the number of arguments is different from 1, the program displays nothing.

### Usage

```console
student@ubuntu:~/rot13/test$ go build
student@ubuntu:~/rot13/test$ ./test "abc"
nop
student@ubuntu:~/rot13/test$ ./test "hello there"
uryyb gurer
student@ubuntu:~/rot13/test$ ./test "HELLO, HELP"
URYYB, URYC
student@ubuntu:~/rot13/test$ ./test
student@ubuntu:~/rot13/test$
```
