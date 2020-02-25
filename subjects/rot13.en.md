## rot13

### Instructions

Write a program that takes a `string` and displays it, replacing each of its
letters by the letter 13 spaces ahead in alphabetical order.

- 'z' becomes 'm' and 'Z' becomes 'M'. Case remains unaffected.

- The output will be followed by a newline (`'\n'`).

- If the number of arguments is different from 1, the program displays a newline (`'\n'`).

### Usage

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test "abc"
nop
student@ubuntu:~/[[ROOT]]/test$ ./test "hello there"
uryyb gurer
student@ubuntu:~/[[ROOT]]/test$ ./test "HELLO, HELP"
URYYB, URYC
student@ubuntu:~/[[ROOT]]/test$ ./test

student@ubuntu:~/[[ROOT]]/test$
```
